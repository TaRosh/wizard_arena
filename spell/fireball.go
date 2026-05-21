package spell

var (
	GestureCircle   = GetTemplate("shied", "templating/a.png")
	GestureY        = GetTemplate("Thunderbalt", "templating/b.png")
	GestureWind     = GetTemplate("Wind", "templating/c.png")
	GestureFireball = GetTemplate("Fireball", "templating/d.png")
)

var Spells []GestureTemplate

func init() {
	Spells = append(Spells, GestureCircle, GestureY, GestureWind, GestureFireball)
}
