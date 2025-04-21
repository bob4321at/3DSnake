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
		return vec4(0,0,0,0)
	}
`
