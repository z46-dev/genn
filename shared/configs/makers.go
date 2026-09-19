package configs

func makeGreenFood(source *Definition) (out *Definition) {
	out = New().
		Label("Green " + *source.Label).
		Parent(source).
		Value(*source.Value * 200).
		Color(1).
		Build()

	*out.Body.Health *= 5
	*out.Body.Damage *= 1.25

	if out.Body.Shield != nil {
		*out.Body.Shield *= 5
	}

	if out.Body.Resist != nil {
		*out.Body.Resist *= 2
	}

	if out.Body.Regeneration != nil {
		*out.Body.Regeneration *= 2
	}

	return
}

type MakeAutoOptions struct {
	name        *string
	def         *Definition
	size        *float64
	independent *bool
}

func NewMakeAutoOptions() (out *MakeAutoOptions) {
	out = &MakeAutoOptions{}
	return
}

func (o *MakeAutoOptions) Name(name string) (self *MakeAutoOptions) {
	o.name = &name
	self = o
	return
}

func (o *MakeAutoOptions) Definition(def *Definition) (self *MakeAutoOptions) {
	o.def = def
	self = o
	return
}

func (o *MakeAutoOptions) Size(size float64) (self *MakeAutoOptions) {
	o.size = &size
	self = o
	return
}

func (o *MakeAutoOptions) Independent(independent bool) (self *MakeAutoOptions) {
	o.independent = &independent
	self = o
	return
}

func makeAuto(def *Definition, opts *MakeAutoOptions) (out *Definition) {
	if opts.name == nil {
		var name string = "Auto-" + *def.Label
		opts.name = &name
	}

	if opts.def == nil {
		opts.def = AutoTurret
	}

	if opts.size == nil {
		var size float64 = 10 * arrasToGennScale
		opts.size = &size
	}

	if opts.independent == nil {
		var independent bool = true
		opts.independent = &independent
	}

	out = cloneDef(opts.def)

	var builder *DefBuilder = New()
	builder.Definition = out

	builder.Label(*opts.name)
	builder.Turret(
		NewTurret(*opts.size, 0, 0, 180, 360, 1).
			Definitions(opts.def, NewDummy().Independent(*opts.independent).Build()).
			Build(),
	)

	out = builder.Build()
	return
}
