// SPDX-License-Identifier: MIT

package macos

import (
	"github.com/issue9/localeutil"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

type Options struct {
	Catalog           catalog.Catalog
	Category          Category // 应用分类
	Copyright         string
	MinOS             string // https://developer.apple.com/documentation/bundleresources/information_property_list/lsminimumsystemversion?language=objc
	Languages         []language.Tag
	ReadableCopyright localeutil.LocaleStringer
	DisplayName       localeutil.LocaleStringer
	GetInfoString     localeutil.LocaleStringer
}

type Category string

// 苹果应用的分类
// https://developer.apple.com/documentation/bundleresources/information_property_list/lsapplicationcategorytype
const (
	CategoryBusiness          = "public.app-category.business"
	CategoryDeveloperTools    = "public.app-category.developer-tools"
	CategoryEducation         = "public.app-category.education"
	CategoryEntertainment     = "public.app-category.entertainment"
	CategoryFinance           = "public.app-category.finance"
	CategoryGames             = "public.app-category.games"
	CategoryActionGames       = "public.app-category.action-games"
	CategoryAdventureGames    = "public.app-category.adventure-games"
	CategoryArcadeGames       = "public.app-category.arcade-games"
	CategoryBoardGames        = "public.app-category.board-games"
	CategoryCardGames         = "public.app-category.card-games"
	CategoryCasinoGames       = "public.app-category.casino-games"
	CategoryDiceGames         = "public.app-category.dice-games"
	CategoryEducationalGames  = "public.app-category.educational-games"
	CategoryFamilyGames       = "public.app-category.family-games"
	CategoryKidsGames         = "public.app-category.kids-games"
	CategoryMusicGames        = "public.app-category.music-games"
	CategoryPuzzleGames       = "public.app-category.puzzle-games"
	CategoryRacingGames       = "public.app-category.racing-games"
	CategoryRolePlayingGames  = "public.app-category.role-playing-games"
	CategorySimulationGames   = "public.app-category.simulation-games"
	CategorySportsGames       = "public.app-category.sports-games"
	CategoryStrategyGames     = "public.app-category.strategy-games"
	CategoryTriviaGames       = "public.app-category.trivia-games"
	CategoryWordGames         = "public.app-category.word-games"
	CategoryGraphicsDesign    = "public.app-category.graphics-design"
	CategoryHealthcareFitness = "public.app-category.healthcare-fitness"
	CategoryLifestyle         = "public.app-category.lifestyle"
	CategoryMedical           = "public.app-category.medical"
	CategoryMusic             = "public.app-category.music"
	CategoryNews              = "public.app-category.news"
	CategoryPhotography       = "public.app-category.photography"
	CategoryProductivity      = "public.app-category.productivity"
	CategoryReference         = "public.app-category.reference"
	CategorySocialNetworking  = "public.app-category.social-networking"
	CategorySports            = "public.app-category.sports"
	CategoryTravel            = "public.app-category.travel"
	CategoryUtilities         = "public.app-category.utilities"
	CategoryVideo             = "public.app-category.video"
	CategoryWeather           = "public.app-category.weather"
)
