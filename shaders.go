package main

var Test_Shader = `//kage:unit pixels
	package main

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)
		below := imageSrc0At(vec2(srcPos.x+(targetCoords.x/64)-(8), srcPos.y+(targetCoords.y/64)-(6)))
		closebelow := imageSrc0At(vec2(srcPos.x+(targetCoords.x/128)-(8), srcPos.y+(targetCoords.y/128)-(6)))

		if col != vec4(0,0,0,0) {
			return vec4(col.x, col.y, col.z, col.w)
		}
		if closebelow.w != 0 {
			return vec4(closebelow.x - 0.05, closebelow.y - 0.05, closebelow.z - 0.05, 1)
		}
		if below.w != 0 {
			return vec4(below.x- 0.1, below.y - 0.1, below.z - 0.1, 1)
		}
	}
`

var Silly_Shader = `//kage:unit pixels
	package main

	var Time float

	func Fragment(targetCoords vec4, srcPos vec2, _ vec4) vec4 {
		col := imageSrc0At(srcPos.xy)
		below := imageSrc0At(vec2(srcPos.x+(targetCoords.x/64), srcPos.y+(targetCoords.y/64)))
		closebelow := imageSrc0At(vec2(srcPos.x+(targetCoords.x/128), srcPos.y+(targetCoords.y/128)))

		// farbelow := imageSrc0At(vec2(srcPos.x-6, srcPos.y))
		// closebelow := imageSrc0At(vec2(srcPos.x, srcPos.y-3))
		// veryfarbelow := imageSrc0At(vec2(srcPos.x-4, srcPos.y-3))
		if col != vec4(0,0,0,0) {
			return vec4(col.x+sin(Time+(srcPos.x/1280)), col.y+cos(Time+(srcPos.y/720)), col.z, col.w)
		}
		if closebelow.w != 0 {
			return vec4(closebelow.x+sin(Time+srcPos.x) - 0.05, closebelow.y+cos(Time+srcPos.y) - 0.05, closebelow.z - 0.05, 1)
		}
		if below.w != 0 {
			return vec4(below.x +sin(Time+srcPos.x)- 0.1, below.y+cos(Time+srcPos.y) - 0.1, below.z - 0.1, 1)
		}

		// if farbelow.w != 0 {
			// return vec4(farbelow.x - 0.1, farbelow.y - 0.1, farbelow.z - 0.1, 1)
		// }
		// if closebelow.w != 0 {
			// return vec4(closebelow.x - 0.1, closebelow.y - 0.1, closebelow.z - 0.1, 1)
		// }
		// if veryfarbelow.w != 0 {
			// return vec4(veryfarbelow.x - 0.1, veryfarbelow.y - 0.1, veryfarbelow.z - 0.1, 1)
		// }
		return vec4(0,0,0,0)
	}
`
