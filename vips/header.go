package vips

// #cgo pkg-config: vips
// #include "header.h"
import "C"

func vipsHasICCProfile(in *C.VipsImage) bool {
	return int(C.has_icc_profile(in)) > 0
}

func vipsTransformICCProfile(in *C.VipsImage) (*C.VipsImage, error) {
	var out *C.VipsImage

	if res := int(C.icc_transform(in, &out)); res != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

func vipsRemoveICCProfile(in *C.VipsImage) bool {
	return fromGboolean(C.remove_icc_profile(in))
}

func vipsRemoveMetadata(in *C.VipsImage) {
	C.remove_metadata(in)
}

func vipsGetMetaOrientation(in *C.VipsImage) int {
	return int(C.get_meta_orientation(in))
}

func vipsSetMetaOrientation(in *C.VipsImage, orientation int) {
	C.set_meta_orientation(in, C.int(orientation))
}