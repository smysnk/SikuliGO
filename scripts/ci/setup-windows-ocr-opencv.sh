#!/usr/bin/env bash
set -euo pipefail

PREFIX="/mingw64"
OPENCV_ROOT="/c/opencv/build/install"
OPENCV_INCLUDE="${OPENCV_ROOT}/include"
OPENCV_LIB="${OPENCV_ROOT}/x64/mingw/lib"
OPENCV_VERSION_SUFFIX="4130"

OPENCV_MODULES=(
  core
  face
  videoio
  imgproc
  highgui
  imgcodecs
  objdetect
  features2d
  video
  dnn
  xfeatures2d
  plot
  tracking
  img_hash
  calib3d
  bgsegm
  photo
  aruco
  wechat_qrcode
  ximgproc
  mcc
)

mkdir -p "${OPENCV_ROOT}/x64/mingw"

if [[ ! -e "${OPENCV_INCLUDE}" ]]; then
  ln -s "${PREFIX}/include/opencv4" "${OPENCV_INCLUDE}"
fi
if [[ ! -e "${OPENCV_LIB}" ]]; then
  ln -s "${PREFIX}/lib" "${OPENCV_LIB}"
fi

# gosseract links with -llept while MSYS2 publishes libleptonica.
if [[ ! -e "${PREFIX}/lib/liblept.dll.a" && -e "${PREFIX}/lib/libleptonica.dll.a" ]]; then
  ln -sf "${PREFIX}/lib/libleptonica.dll.a" "${PREFIX}/lib/liblept.dll.a"
fi

# gocv v0.43.0 hardcodes version-suffixed OpenCV import libs on Windows
# (for example: -lopencv_core4130), while MSYS2 installs versionless
# import libraries (for example: libopencv_core.dll.a). Mirror the names
# gocv expects into the C:/opencv compatibility directory.
for module in "${OPENCV_MODULES[@]}"; do
  src="${PREFIX}/lib/libopencv_${module}.dll.a"
  dst="${OPENCV_LIB}/libopencv_${module}${OPENCV_VERSION_SUFFIX}.dll.a"
  if [[ -f "${src}" && ! -e "${dst}" ]]; then
    ln -sf "${src}" "${dst}"
  fi
done

PKG_CONFIG_PATH_VALUE="${PREFIX}/lib/pkgconfig${PKG_CONFIG_PATH:+:${PKG_CONFIG_PATH}}"
CGO_CFLAGS_VALUE="-I${PREFIX}/include/opencv4 -IC:/opencv/build/install/include"
CGO_CPPFLAGS_VALUE="${CGO_CFLAGS_VALUE}"
CGO_CXXFLAGS_VALUE="${CGO_CFLAGS_VALUE}"
CGO_LDFLAGS_VALUE="-L${PREFIX}/lib -LC:/opencv/build/install/x64/mingw/lib -llept -ltesseract"

if [[ -n "${GITHUB_ENV:-}" ]]; then
  {
    echo "PKG_CONFIG_PATH=${PKG_CONFIG_PATH_VALUE}"
    echo "CGO_CFLAGS=${CGO_CFLAGS_VALUE}"
    echo "CGO_CPPFLAGS=${CGO_CPPFLAGS_VALUE}"
    echo "CGO_CXXFLAGS=${CGO_CXXFLAGS_VALUE}"
    echo "CGO_LDFLAGS=${CGO_LDFLAGS_VALUE}"
    echo "CC=gcc"
    echo "CXX=g++"
    echo "TESSDATA_PREFIX=${PREFIX}/share/tessdata"
  } >> "${GITHUB_ENV}"
fi

if [[ -n "${GITHUB_PATH:-}" ]]; then
  echo "${PREFIX}/bin" >> "${GITHUB_PATH}"
fi

export PKG_CONFIG_PATH="${PKG_CONFIG_PATH_VALUE}"
export CGO_CFLAGS="${CGO_CFLAGS_VALUE}"
export CGO_CPPFLAGS="${CGO_CPPFLAGS_VALUE}"
export CGO_CXXFLAGS="${CGO_CXXFLAGS_VALUE}"
export CGO_LDFLAGS="${CGO_LDFLAGS_VALUE}"
export CC="gcc"
export CXX="g++"
export TESSDATA_PREFIX="${PREFIX}/share/tessdata"

pkg-config --cflags lept tesseract
pkg-config --libs lept tesseract
test -f "${OPENCV_LIB}/libopencv_core${OPENCV_VERSION_SUFFIX}.dll.a"
