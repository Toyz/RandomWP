package settings

import "github.com/Toyz/GoHaven"

type Config struct {
	// Folder to save images too defaults to RandomWP in pictures fodler
	SaveFolder string
	// Save current image to given folder (Defaults: Desktop)
	SaveCurrentImageFolder string
	// cats specifies the enabled wallpaper categories.
	Category GoHaven.Categories
	// purity specifies the enabled purity modes.
	Purity GoHaven.Purity

	// TODO: Actually make this usable one day
	// res specifies the enabled screen resolutions.
	//Resolution wallhaven.Resolutions

	// ratios specifies the enabled aspect rations.
	Ratio GoHaven.Ratios
	// Auto start on boot
	AutoStart bool
	// Send notifications when desktop background changes (Buggy feature)
	Notify bool
	// How many seondds to wait before setting the next wallpaper
	Delay int64
	// Auto delete the image that was set...
	AutoDelete bool
	// Last Image ID
	LastImageID GoHaven.ID

	/* Private */
	confFile string
}
