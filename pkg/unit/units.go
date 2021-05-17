package unit

import (
	"math/big"

	"github.com/nkanaev/numb/pkg/consts"
)

var one = big.NewRat(1, 1)

func div(a *big.Rat, x int64) *big.Rat {
	num := new(big.Rat).Set(a)
	num.Quo(num, big.NewRat(x, 1))
	return num
}

func num(x string) *big.Rat {
	rat, ok := new(big.Rat).SetString(x)
	if !ok {
		panic("unable to parse: " + x)
	}
	return rat
}

var units = []baseUnit{
	{d: LENGTH, name: "m", long: "meter, metre", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{d: LENGTH, name: "in", long: "inch", value: num("0.0254")},
	{d: LENGTH, name: "ft", long: "foot, feet", value: num("0.3048")},
	{d: LENGTH, name: "yd", long: "yard, yards", value: num("0.9144")},
	{d: LENGTH, name: "mi", long: "mile, miles", value: num("1609.344")},

	{d: LENGTH, name: "li", long: "link", value: num("0.201168")},
	{d: LENGTH, name: "rd", long: "rod", value: num("5.0292")},
	{d: LENGTH, name: "ch", long: "chain", value: num("20.1")},
	{d: LENGTH, name: "angstrom", value: exp(10, -10)},
	{d: LENGTH, name: "mil", value: num("0.0000254")},
	{d: LENGTH, name: "au", long: "astronomical-unit", value: num("149597870700"), info: "accepted for use with the SI"},
	{d: LENGTH, name: "ly", long: "lightyeaar, light-year", value: num("9460730472580800"), prefixes: &metricPrefixes, info: "accepted for use with the SI"},
	{d: LENGTH, name: "lightsecond", long: "lightsecond, light-second", value: num("299792458"), prefixes: &metricPrefixes},

	{d: TIME, name: "s", long: "sec, second", value: num("1"), prefixes: &metricPrefixes, info: "SI base unit"},
	{d: TIME, name: "min", long: "minute", value: num("60"), info: "accepted for use with the SI (1 min = 60 s)"},
	{d: TIME, name: "h", long: "hr, hour", value: num("3600"), info: "accepted for use with the SI (1 h = 60 min)"},
	{d: TIME, name: "d", long: "day", value: num("86400"), info: "accepted for use with the SI (1 day = 24 h)"},
	{d: TIME, name: "week", value: num("604800")},
	{d: TIME, name: "month", value: num("2629800"), info: "1/12th of Julian Year"},
	{d: TIME, name: "year", value: num("31557600"), info: "Julian Year (365.25 days)"},
	{d: TIME, name: "decade", value: num("315576000"), info: "Julian decade"},
	{d: TIME, name: "century", value: num("3155760000"), info: "Julian century"},
	{d: TIME, name: "millenium", value: num("31557600000"), info: "Julian millenium"},

	{d: MASS, name: "g", long: "gram", value: num("0.001"), prefixes: &metricPrefixes, info: "(0.001 kg - SI base unit)"},
	{d: MASS, name: "t", long: "tonne, metric-ton", value: num("1000"), prefixes: &metricPrefixes, info: "accepted for use with the SI (1 t = 1000 kg)"},
	{d: MASS, name: "Da", long: "dalton", value: num("1.6605402e-27"), info: "accepted for use with the SI"},

	{d: TEMPERATURE, name: "K", long: "kelvin", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{d: TEMPERATURE, name: "°C, degC", long: "celsius", value: one, offset: num("273.15"), info: "SI derived unit"},
	{d: TEMPERATURE, name: "°F, degF", long: "fahrenheit", value: num("10/18"), offset: num("459.67")},

	{d: ANGLE, name: "rad", long: "radian", value: one, info: "SI derived unit"},
	{d: ANGLE, name: "°", long: "deg, degree", value: div(consts.PI, 180), info: "accepted for use with the SI"},
	{d: ANGLE, name: "arcsec", value: div(consts.PI, 648000), info: "accepted for use with the SI (pi / 648000)"},
	{d: ANGLE, name: "arcmin", value: div(consts.PI, 10800), info: "accepted for use with the SI (pi / 10800)"},
	{d: ANGLE, name: "grad", long: "grade, gradian", value: div(consts.PI, 200)},
	{d: ANGLE, name: "cycle", value: div(consts.PI, 2)},

	{d: DIGITAL, name: "bit", long: "bit", value: one, prefixes: &digitalPrefixes},
	{d: DIGITAL, name: "B", long: "byte", value: num("8"), prefixes: &digitalPrefixes},

	{d: AREA, name: "m², m2", value: one, prefixes: &metricPrefixes, prefixpow: 2},
	{d: AREA, name: "in², in2, sqin", value: num("0.00064516")},
	{d: AREA, name: "ft², ft2, sqft", long: "sqfeet", value: num("0.09290304")},
	{d: AREA, name: "yd², sqyd", long: "sqyard", value: num("0.83612736")},
	{d: AREA, name: "rd², rd2, sqrd", value: num("25.29295")},
	{d: AREA, name: "ch², sqch", value: num("404.6873")},
	{d: AREA, name: "mi², sqmil, sqmi", long: "sqmile", value: num("6.4516e-10")},
	{d: AREA, name: "acre", value: num("4046.86")},
	{d: AREA, name: "ha", long: "hectare", value: num("10000"), info: "accepted for use with the SI (1 ha = 10,000 m²)"},

	{d: VOLUME, name: "m³, m3", value: one, prefixes: &metricPrefixes, prefixpow: 3},
	{d: VOLUME, name: "l, lt", long: "liter, litre", value: num("0.001"), prefixes: &metricPrefixes, info: "accepted for use with the SI (1 l = 0.001 m³)"},
	{d: VOLUME, name: "in³, in3, cuin", value: num("1.6387064e-5")},
	{d: VOLUME, name: "ft³, ft3, cuft", value: num("0.028316846592")},
	{d: VOLUME, name: "yd³, yd3, cuyd", value: num("0.764554857984")},
	{d: VOLUME, name: "teaspoon", value: num("0.000005")},
	{d: VOLUME, name: "tablespoon", value: num("0.000015")},

	{d: ENERGY, name: "J", long: "joule", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ENERGY, name: "Wh", long: "watt-hour", value: num("3600"), prefixes: &metricPrefixes},
	{d: ENERGY, name: "eV", long: "electronvolt", value: num("1.602176565e-19"), prefixes: &metricPrefixes, info: "accepted for use with the SI"},
	{d: ENERGY, name: "erg", value: exp(10, -7)},

	{d: FORCE, name: "N", long: "newton", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: FORCE, name: "dyn", long: "dyne", value: exp(10, -5)},
	{d: FORCE, name: "lbf", long: "poundforce", value: num("4.4482216152605")},
	{d: FORCE, name: "kip", value: num("4448.2216")},
	{d: FORCE, name: "pdl", long: "poundal", value: num("0.138254954376")},

	{d: PRESSURE, name: "Pa", long: "pascal", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: PRESSURE, name: "psi", value: num("6894.757"), info: "US/Imperial unit (pound per square inch)"},
	{d: PRESSURE, name: "at", long: "technical-atmosphere", value: num("98066.5")},
	{d: PRESSURE, name: "atm", long: "atmosphere, standard-atmosphere", value: num("101325")},
	{d: PRESSURE, name: "bar", long: "bar", value: num("100000"), prefixes: &metricPrefixes},
	{d: PRESSURE, name: "torr", long: "Torr", value: num("101325/760")},
	{d: PRESSURE, name: "mmHg", value: num("133.322387415")},
	{d: PRESSURE, name: "mmH2O", value: num("9.80665")},
	{d: PRESSURE, name: "cmH2O", value: num("98.0665")},

	{d: POWER, name: "W", long: "watt", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: POWER, name: "hp", long: "horsepower", value: num("745.69987158227"), info: "mechanical horsepower"},

	{d: RADIOACTIVITY, name: "Bq", long: "becquierel", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: RADIOACTIVITY, name: "Ci", long: "curie", value: num("3.7e10")},
	{d: RADIOACTIVITY, name: "Rd", long: "rutherford", value: exp(10, 6)},

	{d: AMOUNT_OF_SUBSTANCE, name: "mol", long: "mole", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{d: CATALYCTIC_ACTIVITY, name: "kat", long: "katal", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ELECTRIC_CAPACITANCE, name: "F", long: "farad", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ELECTRIC_CHARGE, name: "C", long: "coulomb", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ELECTRIC_CONDUCTANCE, name: "S", long: "siemens", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ELECTRIC_CURRENT, name: "A, amp", long: "ampere", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{d: ELECTRIC_INDUCTANCE, name: "H", long: "henry", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ELECTRIC_POTENTIAL, name: "V", long: "volt", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ELECTRIC_RESISTANCE, name: "Ω, ohm", long: "ohm", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: FREQUENCY, name: "Hz", long: "hertz", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: ILLUMINANCE, name: "lx", long: "lux", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: IONIZING_RADIATION, name: "Sv", long: "sievert", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: LUMINOUS_FLUX, name: "lm", long: "lumen", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: LUMINOUS_INTENSITY, name: "cd", long: "candela", value: one, prefixes: &metricPrefixes, info: "SI base unit"},
	{d: MAGNETIC_FLUX, name: "Wb", long: "weber", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: MAGNETIC_FLUX_DENSITY, name: "T", long: "tesla", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: RADIATION_DOSE, name: "Gy", long: "gray", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
	{d: SOLID_ANGLE, name: "sr", long: "steradian", value: one, prefixes: &metricPrefixes, info: "SI derived unit"},
}
