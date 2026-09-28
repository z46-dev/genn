package game

import (
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"strings"

	"github.com/z46-dev/gamelib/hshg"
	"github.com/z46-dev/gamelib/vector"
	"github.com/z46-dev/genn/shared"
	"github.com/z46-dev/genn/shared/configs"
)

func NewRoom(c *Configuration) (r *Room, err error) {
	r = &Room{
		C: c,
	}

	r.Zones.Type = make(map[configs.RoomCellType][]*Zone)

	r.Resize(c.Width, c.Height)
	err = r.SetCells(c.Setup, c.XGrid, c.YGrid)
	return
}

// Resize the room to a new width and height.
func (r *Room) Resize(width, height float64) {
	r.Width, r.Height = width, height
	r.HalfWidth, r.HalfHeight = width/2, height/2
	r.Scale.Square = width * height / 100000000
	r.Scale.Linear = math.Sqrt(r.Scale.Square)
	r.MaxFood = width * height / 100000 * r.C.FoodAmount
	r.NeedsBroadcast = true
}

// SetCells sets the room's cell setup and grid dimensions.
// It also indexes the zones for each cell type.
func (r *Room) SetCells(setup [][]string, xGrid, yGrid int) (err error) {
	if len(setup) != yGrid || slices.ContainsFunc(setup, func(x []string) (incorrect bool) {
		incorrect = len(x) != xGrid
		return
	}) {
		err = fmt.Errorf("room setup does not match grid dimensions: %dx%d", yGrid, xGrid)
		return
	}

	r.Setup = make([][]configs.RoomCellType, yGrid)
	r.Zones.Location = make([][]*Zone, yGrid)
	for y := range yGrid {
		r.Setup[y] = make([]configs.RoomCellType, xGrid)
		r.Zones.Location[y] = make([]*Zone, xGrid)
		for x := range xGrid {
			r.Setup[y][x] = configs.RoomCellTypeIDs[setup[y][x]]
		}
	}

	r.Grid.X, r.Grid.Y = xGrid, yGrid

	for cell := range configs.RoomCellTypeSENTINEL {
		r.IndexType(cell)
	}

	r.NeedsBroadcast = true
	return
}

// SetCell sets a specific cell in the room's setup and indexes the zones for that cell type.
func (r *Room) SetCell(x, y int, cell configs.RoomCellType) {
	if x < 0 || x >= r.Grid.X || y < 0 || y >= r.Grid.Y {
		return
	}

	var oldCell configs.RoomCellType = r.Setup[y][x]
	r.Setup[y][x] = cell

	r.IndexType(oldCell)
	r.IndexType(cell)
	r.NeedsBroadcast = true
}

// IndexType indexes the zones for a specific cell type in the room.
func (r *Room) IndexType(cell configs.RoomCellType) {
	// Since this should run rarely and is not a hot path, we can afford to be a bit lazy here and use string manipulation
	var (
		stringCell       string = configs.RoomCellTypeNames[cell]
		killZone, portal bool
		team             int
	)

	switch {
	case slices.ContainsFunc([]string{"bas", "bap", "bad"}, func(prefix string) (match bool) {
		match = strings.HasPrefix(stringCell, prefix)
		return
	}):
		killZone = true
		team = int(stringCell[len(stringCell)-1] - '0')
	case strings.HasPrefix(stringCell, "dom"):
		team = int(stringCell[len(stringCell)-1] - '0')
	case strings.HasPrefix(stringCell, "ptl"):
		portal = true
		team = int(stringCell[len(stringCell)-1] - '0')
	}

	r.Zones.Type[cell] = r.Zones.Type[cell][:0]
	for y := range r.Setup {
		for x := range r.Setup[y] {
			if r.Setup[y][x] == cell {
				var zone *Zone = &Zone{
					GridX: x,
					GridY: y,
					Type:  cell,
					AABB2: &hshg.AABB2[float64]{
						X1: float64(x) * r.Width / float64(r.Grid.X),
						Y1: float64(y) * r.Height / float64(r.Grid.Y),
						X2: float64(x+1) * r.Width / float64(r.Grid.X),
						Y2: float64(y+1) * r.Height / float64(r.Grid.Y),
					},
					KillZone: killZone,
					Portal:   portal,
					Team:     team,
				}

				r.Zones.Type[cell] = append(r.Zones.Type[cell], zone)
				r.Zones.Location[y][x] = zone
			}
		}
	}

	if cell == configs.RoomCellTypeNest {
		r.NestFoodAmount = 1.5 * math.Sqrt(float64(len(r.Zones.Type[cell]))) / float64(r.Grid.X) / float64(r.Grid.Y)
	}
}

// IsInRoom checks if a given position is within the room's boundaries.
func (r *Room) IsInRoom(pos *vector.Vec2[float64]) (in bool) {
	in = pos.X >= -r.HalfWidth && pos.X <= r.HalfWidth && pos.Y >= -r.HalfHeight && pos.Y <= r.HalfHeight
	return
}

// Gets a random position within the room.
func (r *Room) Random() (pos *vector.Vec2[float64]) {
	pos = vector.NewVec2(-r.HalfWidth+rand.Float64()*r.Width, -r.HalfHeight+rand.Float64()*r.Height)
	return
}

// RandomType gets a random position within the room for a specific cell type.
// If there are no zones for that cell type, it returns a random position within the room.
func (r *Room) RandomType(cell configs.RoomCellType) (pos *vector.Vec2[float64]) {
	if len(r.Zones.Type[cell]) == 0 {
		pos = r.Random()
		return
	}

	var zone *Zone = shared.Choose(r.Zones.Type[cell])
	pos = vector.NewVec2(zone.X1+rand.Float64()*(zone.X2-zone.X1), zone.Y1+rand.Float64()*(zone.Y2-zone.Y1))
	return
}

// Gauss returns a random position within the room, with a clustering factor that determines how tightly the values are clustered around the center of the room.
func (r *Room) Gauss(clustering float64) (pos *vector.Vec2[float64]) {
	pos = vector.NewVec2(shared.Gauss(r.HalfWidth, r.Width/clustering), shared.Gauss(r.HalfHeight, r.Height/clustering))
	if !r.IsInRoom(pos) {
		pos.X, pos.Y = shared.Clamp(pos.X, -r.HalfWidth, r.HalfWidth), shared.Clamp(pos.Y, -r.HalfHeight, r.HalfHeight)
	}

	return
}

// GaussInverse returns a random position within the room, with a clustering factor that determines how tightly the values are clustered around the edges of the room.
func (r *Room) GaussInverse(clustering float64) (pos *vector.Vec2[float64]) {
	pos = vector.NewVec2(shared.GaussInverse(-r.HalfWidth, r.HalfWidth, clustering), shared.GaussInverse(-r.HalfHeight, r.HalfHeight, clustering))
	if !r.IsInRoom(pos) {
		pos.X, pos.Y = shared.Clamp(pos.X, -r.HalfWidth, r.HalfWidth), shared.Clamp(pos.Y, -r.HalfHeight, r.HalfHeight)
	}

	return
}

// GaussRing returns a random position within the room, in a ring with a given radius and clustering factor.
func (r *Room) GaussRing(radius, clustering float64) (pos *vector.Vec2[float64]) {
	pos = shared.GaussRing(radius, clustering)
	pos.X, pos.Y = pos.X+r.HalfWidth, pos.Y+r.HalfHeight

	if !r.IsInRoom(pos) {
		pos.X, pos.Y = shared.Clamp(pos.X, -r.HalfWidth, r.HalfWidth), shared.Clamp(pos.Y, -r.HalfHeight, r.HalfHeight)
	}

	return
}

// Location gets the zone at a given position in the room. If the position is outside the room, it returns nil.
func (r *Room) Location(pos *vector.Vec2[float64]) (loc *Zone) {
	if !r.IsInRoom(pos) {
		return
	}

	var idx, idy int = int((pos.X + r.HalfWidth) / r.Width * float64(r.Grid.X)), int((pos.Y + r.HalfHeight) / r.Height * float64(r.Grid.Y))
	if idx < 0 || idx >= r.Grid.X || idy < 0 || idy >= r.Grid.Y {
		return
	}

	loc = r.Zones.Location[idy][idx]
	return
}

// IsIn checks if a given position is within a specific cell type in the room.
// If the position is outside the room, it returns false.
func (r *Room) IsIn(cell configs.RoomCellType, location *vector.Vec2[float64]) (in bool) {
	if in = !r.IsInRoom(location); !in {
		return
	}

	var loc *Zone = r.Location(location)
	in = loc != nil && loc.Type == cell
	return
}

// IsInSafe checks if a given position is within a safe zone in the room.
// A safe zone is defined as a zone that is not a kill zone, not a portal, and not a nest.
// If the position is outside the room, it returns false.
func (r *Room) IsInSafe(location *vector.Vec2[float64]) (in bool) {
	if in = !r.IsInRoom(location); !in {
		return
	}

	var loc *Zone = r.Location(location)
	in = loc != nil && !loc.KillZone && !loc.Portal && loc.Type != configs.RoomCellTypeNest
	return
}

func (r *Room) GaussType(cell configs.RoomCellType, clustering float64) (pos *vector.Vec2[float64]) {
	if len(r.Zones.Type[cell]) == 0 {
		pos = r.Gauss(clustering)
		return
	}

	var zone *Zone = shared.Choose(r.Zones.Type[cell])
	pos = vector.NewVec2(shared.Gauss(zone.X2-zone.X1, r.Width/float64(r.Grid.X)/clustering), shared.Gauss(zone.Y2-zone.Y1, r.Height/float64(r.Grid.Y)/clustering))

	if !r.IsInRoom(pos) {
		pos.X, pos.Y = shared.Clamp(pos.X, -r.HalfWidth, r.HalfWidth), shared.Clamp(pos.Y, -r.HalfHeight, r.HalfHeight)
	}

	return
}
