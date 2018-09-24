package rest

import (
	"net/http"
)

var routes = ConfiguredRoutes{
	Route{
		Name:        "GetNearLocation",
		Method:      "GET",
		Pattern:     "/near",
		HandlerFunc: HandleGetNear},
	Route{
		Name:        "GetByState",
		Method:      "GET",
		Pattern:     "/state/{state-code}",
		HandlerFunc: HandleGetState},
	Route{
		Name:        "GetAllProducts",
		Method:      "GET",
		Pattern:     "/product",
		HandlerFunc: HandleGetAllProduct},
	Route{
		Name:        "GetHybridProducts",
		Method:      "GET",
		Pattern:     "/product/hybrid",
		HandlerFunc: HandleGetHybridProduct},
	Route{
		Name:        "GetSativaProducts",
		Method:      "GET",
		Pattern:     "/product/sativa",
		HandlerFunc: HandleGetSativaProduct},
	Route{
		Name:        "GetIndicaProducts",
		Method:      "GET",
		Pattern:     "/product/indica",
		HandlerFunc: HandleGetIndicaProduct},
	Route{
		Name:        "GetOtherProducts",
		Method:      "GET",
		Pattern:     "/other-product",
		HandlerFunc: HandleGetOtherProduct},
	Route{
		Name:        "GetOtherProductsEdible",
		Method:      "GET",
		Pattern:     "/other-product/edible",
		HandlerFunc: HandleGetOtherProductEdible},
	Route{
		Name:        "GetOtherProductsConcentrate",
		Method:      "GET",
		Pattern:     "/other-product/concentrate",
		HandlerFunc: HandleGetOtherProductConcentrate},
	Route{
		Name:        "GetOtherProductsConcentrateCartridge",
		Method:      "GET",
		Pattern:     "/other-product/concentrate/cartridge",
		HandlerFunc: HandleGetOtherProductConcentrateCartridge},
	Route{
		Name:        "GetOtherProductsConcentrateIngestible",
		Method:      "GET",
		Pattern:     "/other-product/concentrate/ingestible",
		HandlerFunc: HandleGetOtherProductConcentrateIngestible},
	Route{
		Name:        "GetOtherProductsConcentrateSolvent",
		Method:      "GET",
		Pattern:     "/other-product/concentrate/solvent",
		HandlerFunc: HandleGetOtherProductConcentrateSolvent},
	Route{
		Name:        "GetOtherProductsConcentrateSolventless",
		Method:      "GET",
		Pattern:     "/other-product/concentrate/solventless",
		HandlerFunc: HandleGetOtherProductConcentrateSolventless},
	Route{
		Name:        "GetOtherProductsConcentrateTerpene",
		Method:      "GET",
		Pattern:     "/other-product/concentrate/terpene",
		HandlerFunc: HandleGetOtherProductConcentrateTerpene},
	Route{
		Name:        "GetOtherProductsEdibleBeverage",
		Method:      "GET",
		Pattern:     "/other-product/edible/beverage",
		HandlerFunc: HandleGetOtherProductEdibleBeverage},
	Route{
		Name:        "GetOtherProductsEdibleBreakfast",
		Method:      "GET",
		Pattern:     "/other-product/edible/breakfast",
		HandlerFunc: HandleGetOtherProductEdibleBreakfast},
	Route{
		Name:        "GetOtherProductsEdibleBrownie",
		Method:      "GET",
		Pattern:     "/other-product/edible/brownie",
		HandlerFunc: HandleGetOtherProductEdibleBrownie},
	Route{
		Name:        "GetOtherProductsEdibleCandy",
		Method:      "GET",
		Pattern:     "/other-product/edible/candy",
		HandlerFunc: HandleGetOtherProductEdibleCandy},
	Route{
		Name:        "GetOtherProductsEdibleCapsule",
		Method:      "GET",
		Pattern:     "/other-product/edible/capsule",
		HandlerFunc: HandleGetOtherProductEdibleCapsule},
	Route{
		Name:        "GetOtherProductsEdibleChocolate",
		Method:      "GET",
		Pattern:     "/other-product/edible/chocolate",
		HandlerFunc: HandleGetOtherProductEdibleChocolate},
	Route{
		Name:        "GetOtherProductsEdibleCondiment",
		Method:      "GET",
		Pattern:     "/other-product/edible/condiment",
		HandlerFunc: HandleGetOtherProductEdibleCondiment},
	Route{
		Name:        "GetOtherProductsEdibleCookie",
		Method:      "GET",
		Pattern:     "/other-product/edible/cookie",
		HandlerFunc: HandleGetOtherProductEdibleCookie},
	Route{
		Name:        "GetOtherProductsEdibleCooking",
		Method:      "GET",
		Pattern:     "/other-product/edible/cooking",
		HandlerFunc: HandleGetOtherProductEdibleCooking},
	Route{
		Name:        "GetOtherProductsEdibleFrozen",
		Method:      "GET",
		Pattern:     "/other-product/edible/frozen",
		HandlerFunc: HandleGetOtherProductEdibleFrozen},
	Route{
		Name:        "GetOtherProductsEdibleSnack",
		Method:      "GET",
		Pattern:     "/other-product/edible/snack",
		HandlerFunc: HandleGetOtherProductEdibleSnack},
	Route{
		Name:        "GetOtherProductsEdibleTincture",
		Method:      "GET",
		Pattern:     "/other-product/edible/tincture",
		HandlerFunc: HandleGetOtherProductEdibleTincture},
}

type ConfiguredRoutes []Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc AuthHandler
}

type AuthHandler func(http.ResponseWriter, *http.Request)

func (fn AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// TODO: implement authentication

	fn(w, r)
}
