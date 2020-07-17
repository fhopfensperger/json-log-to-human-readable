#!/usr/bin/env bash

# Copyright Florian Hopfensperger <f.hopfensperger@gmail.com.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# The install script is based off of the MIT-licensed script from glide,
# the package manager for Go: https://github.com/Masterminds/glide.sh/blob/master/get
# and the Helm install script https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
: ${INSTALL_DIR:="/usr/local/bin"}
: ${BINARY_NAME:="json-log-to-human-readable"}

installNotes() {
  echo "Installing ${BINARY_NAME}..."
}


# initArch discovers the architecture for this system.
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
    armv5*) ARCH="armv5";;
    armv6*) ARCH="armv6";;
    armv7*) ARCH="arm";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
  echo ARCH: $ARCH
}

# initOS discovers the operating system for this system.
initOS() {
  OS=$(echo `uname`|tr '[:upper:]' '[:lower:]')

  case "$OS" in
    # Minimalist GNU for Windows
    mingw*) OS='windows';;
  esac
  echo OS: $OS
}

# runs the given command as root (detects if we are root already)
runAsRoot() {
  local CMD="$*"

  if [ $EUID -ne 0 ]; then
    CMD="sudo $CMD"
  fi

  $CMD
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  local supported="darwin-386\ndarwin-amd64\nlinux-386\nlinux-amd64\nlinux-arm\nlinux-arm64\nlinux-ppc64le\nlinux-s390x\nwindows-386\nwindows-amd64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuilt binary for ${OS}-${ARCH}."
    echo "To build from source, go to https://github.com/fhopfensperger/json-log-to-human-readable"
    exit 1
  fi

  if ! type "curl" > /dev/null; then
    echo "curl is required"
    exit 1
  fi
}

# checkDesiredVersion checks if the desired version is available.
checkDesiredVersion() {
  if [ "x$DESIRED_VERSION" == "x" ]; then
    # Get tag from release URL
    TAG=$(curl --silent "https://github.com/fhopfensperger/json-log-to-human-readable/releases/latest" | sed 's#.*tag/\(.*\)\".*#\1#')
  else
    TAG=$DESIRED_VERSION
  fi
}

# checkjl2hrInstalledVersion checks which version of json-log-to-human-readable is installed and
# if it needs to be changed.
checkjl2hrInstalledVersion() {
  if [[ -f "$INSTALL_DIR/$BINARY_NAME" ]]; then
    local version=$("$INSTALL_DIR/$BINARY_NAME" -v)
    if [[ "$version" == "$TAG" ]]; then
      echo "${BINARY_NAME} ${version} is already ${DESIRED_VERSION:-latest}"
      return 0
    else
      echo "${BINARY_NAME} ${TAG} is available. Changing from version ${version}."
      return 1
    fi
  else
    return 1
  fi
}

downloadAndInstall() {

  downloadUrl=$(curl -s https://api.github.com/repos/fhopfensperger/json-log-to-human-readable/releases/latest | grep browser_download_url | grep ${OS}_${ARCH} | cut -d ":" -f 2,3 | cut -d " " -f 2,3 | tr -d \")

  echo "downloading.."
  curl -L -s ${downloadUrl} --output /tmp/${BINARY_NAME}.tar.gz

  tar xzf /tmp/${BINARY_NAME}.tar.gz

  echo "copy ${BINARY_NAME} to ${INSTALL_DIR}"
  # runAsRoot mv ${BINARY_NAME} ${INSTALL_DIR}/
  runAsRoot chmod +x /usr/local/bin/${BINARY_NAME}

  location=$(which $BINARY_NAME)
  echo "${BINARY_NAME} binary location: $location"

  version="$($BINARY_NAME -v)"
  echo "${BINARY_NAME} binary version: $version"

}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
      echo "Failed to install $BINARY_NAME"
  fi
  exit $result
}

#Stop execution on any error
trap "fail_trap" EXIT
set -e

installNotes
initArch
initOS
verifySupported
checkDesiredVersion
if ! checkjl2hrInstalledVersion; then
  downloadAndInstall
fi