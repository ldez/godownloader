package main

const shellfn = `

cat /dev/null << EOF
------------------------------------------------------------------------
https://github.com/client9/posixshell - portable posix shell functions
Public domain - http://unlicense.org
https://github.com/client9/posixshell/blob/master/LICENSE.md
but credits (and pull requests) appreciated.
------------------------------------------------------------------------
EOF
is_command() {
  type $1 > /dev/null 2> /dev/null
}
uname_arch() {
  local arch=$(uname -m)
  case $arch in
    x86_64) arch="amd64" ;;
    x86)    arch="386" ;;
    i686)   arch="386" ;;
    i386)   arch="386" ;;
  esac
  echo ${arch}
}
uname_os() {
  local os=$(uname -s | tr '[:upper:]' '[:lower:]')
  echo ${os}
}
untar() {
  TARBALL=$1
  case ${TARBALL} in
  *.tar.gz|*.tgz) tar -xzf ${TARBALL} ;;
  *.tar) tar -xf ${TARBALL} ;;
  *.zip) unzip ${TARBALL} ;;
  *)
    echo "Unknown archive format for ${TARBALL}"
    return 1
  esac
}
mktmpdir() {
   test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"
   mkdir -p ${TMPDIR}
}
http_download() {
  DEST=$1
  SOURCE=$2
  HEADER=$3
  if is_command curl; then
    WGET="curl --fail -sSL"
    test -z "${HEADER}" || WGET="${WGET} -H \"${HEADER}\""
    if [ "${DEST}" != "-" ]; then
      WGET="$WGET -o $DEST"
    fi
  elif is_command wget &> /dev/null; then
    WGET="wget -q -O $DEST"
    test -z "${HEADER}" || WGET="${WGET} --header \"${HEADER}\""
  else
    echo "Unable to find wget or curl.  Exit"
    exit 1
  fi
  if [ "${DEST}" != "-" ]; then
    rm -f "${DEST}"
  fi
  ${WGET} ${SOURCE}
}
github_api() {
  DEST=$1
  SOURCE=$2
  HEADER=""
  case $SOURCE in
  https://api.github.com*)
     test -z "$GITHUB_TOKEN" || HEADER="Authorization: token $GITHUB_TOKEN"
     ;;
  esac
  http_download $DEST $SOURCE $HEADER
}
hash_sha256() {
  TARGET=${1:-$(</dev/stdin)};
  if is_command gsha256sum; then
    gsha256sum $TARGET | cut -d ' ' -f 1
  elif is_command sha256sum; then
    sha256sum $TARGET | cut -d ' ' -f 1
  elif is_command shasum; then
    shasum -a 256 $TARGET | cut -d ' ' -f 1
  elif is_command openssl; then
    openssl -dst openssl dgst -sha256 $TARGET | cut -d ' ' -f a
  else
    echo "Unable to compute hash. exiting"
    exit 1
  fi
}
hash_sha256_verify() {
  TARGET=$1
  SUMS=$2
  BASENAME=${TARGET##*/}
  WANT=$(grep ${BASENAME} ${SUMS} | tr '\t' ' ' | cut -d ' ' -f 1)
  GOT=$(hash_sha256 $TARGET)
  if [ "$GOT" != "$WANT" ]; then
     echo "Checksum for $TARGET did not verify"
     echo "WANT: ${WANT}"
     echo "GOT : ${GOT}"
     exit 1
  fi
}
cat /dev/null << EOF
------------------------------------------------------------------------
End of functions from https://github.com/client9/posixshell 
------------------------------------------------------------------------
EOF
`