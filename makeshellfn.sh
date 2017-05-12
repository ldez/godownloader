#!/bin/sh
set -e

git_clone_or_update() {
  URL=$1
  REPO=${URL##*/}   # foo.git
  REPO=${REPO%.git} # foo
  if [ ! -d "$REPO" ]; then
    git clone ${URL} 
  else
    (cd ${REPO} && git pull > /dev/null)
  fi
}

git_clone_or_update https://github.com/client9/posixshell.git 
cd posixshell

echo "package main"
echo ""
echo 'const shellfn = `'
cat \
  license.sh \
  is_command.sh \
  uname.sh \
  untar.sh \
  mktmpdir.sh \
  http_download.sh \
  github_api.sh \
  hash_sha256.sh \
  license_end.sh | \
  grep -v '^#' |grep -v ' #' | tr -s '\n'

echo '`'
