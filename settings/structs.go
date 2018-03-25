package settings

import "github.com/Toyz/RandomWP/wallhaven"

type Config struct {
	// Folder to save images too defaults to RandomWP in pictures fodler
	SaveFolder string
	// cats specifies the enabled wallpaper categories.
	Category wallhaven.Categories
	// purity specifies the enabled purity modes.
	Purity wallhaven.Purity
	// res specifies the enabled screen resolutions.
	//Resolution wallhaven.Resolutions // TODO : Actually make this usable one day
	// ratios specifies the enabled aspect rations.
	Ratio wallhaven.Ratios
	// Send notifications when desktop background changes (Buggy feature)
	Notify bool
	// How many seondds to wait before setting the next wallpaper
	Delay int64
	// Auto delete the image that was set...
	AutoDelete bool

	/* Private */
	confFile string
}