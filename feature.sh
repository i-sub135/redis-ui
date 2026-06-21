#!/bin/bash

# feature.sh - Command builder untuk membuat fitur baru dari template _example
# Mendukung struktur bersarang: public/auth/validate_request
# Usage: ./feature.sh <scope>/[<domain>/]<feature_name>
# Example: ./feature.sh public/get_user
#          ./feature.sh private/create_session
#          ./feature.sh public/auth/validate_request

set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <scope>/[<domain>/]<feature_name>"
  echo "Example: $0 public/get_user"
  echo "         $0 private/create_session"
  echo "         $0 public/auth/validate_request"
  exit 1
fi

INPUT="$1"

# Validasi format input: harus diawali dengan public/ atau private/
if [[ ! "$INPUT" =~ ^(public|private)/.+ ]]; then
  echo "Error: Format harus <scope>/[<domain>/]<feature_name>"
  echo "Scope harus 'public' atau 'private'"
  echo "Contoh: public/get_user, public/auth/validate_request"
  exit 1
fi

SCOPE=$(echo "$INPUT" | cut -d'/' -f1)
REST=$(echo "$INPUT" | cut -d'/' -f2-)

# Validasi sisa bagian: hanya boleh huruf kecil, underscore, dan slash sebagai pemisah
if [[ ! "$REST" =~ ^[a-z_]+(/[a-z_]+)*$ ]]; then
  echo "Error: Nama fitur dan domain hanya boleh huruf kecil, underscore, dan slash sebagai pemisah"
  echo "Contoh yang valid: get_user, auth/validate_request, admin/user_management"
  exit 1
fi

# Jika tidak ada slash di REST, maka itu adalah fitur langsung di bawah scope
# Jika ada slash, maka bagian terakhir adalah nama fitur, sisanya adalah domain
if [[ "$REST" == */* ]]; then
  # Ada domain: misal auth/validate_request -> domain=auth, feature=validate_request
  DOMAIN_PATH=$(echo "$REST" | sed 's|/[^/]*$||')           # semua kecuali terakhir
  FEATURE_NAME=$(echo "$REST" | awk -F'/' '{print $NF}')   # bagian terakhir
  FULL_PATH="$SCOPE/$DOMAIN_PATH/$FEATURE_NAME"
else
  # Tidak ada domain: misal get_user -> langsung di bawah scope
  DOMAIN_PATH=""
  FEATURE_NAME="$REST"
  FULL_PATH="$SCOPE/$FEATURE_NAME"
fi

# Konversi nama feature ke format yang diperlukan
FEATURE_CAMEL=$(echo "$FEATURE_NAME" | sed -r 's/(^|_)([a-z])/\U\2/g')  # get_user -> GetUser
FEATURE_LOWER="$FEATURE_NAME"                                         # get_user
FEATURE_UPPER=$(echo "$FEATURE_NAME" | tr '[:lower:]' '[:upper:]')    # GET_USER

FEATURE_PATH="source/feature/$FULL_PATH"
TEMPLATE_PATH="source/feature/_example"

echo "Membuat fitur baru: $INPUT"
echo "Lokasi: $FEATURE_PATH"

if [[ -e "$FEATURE_PATH" ]]; then
  echo "Error: Feature path sudah ada: $FEATURE_PATH"
  echo "Hapus atau rename folder tersebut sebelum generate ulang."
  exit 1
fi

# Buat direktori feature dan subdirektori body
mkdir -p "$FEATURE_PATH/body"

# Fungsi untuk menyalin dan mengganti placeholder
copy_and_replace() {
  local src="$1"
  local dst="$2"

  if [[ ! -f "$src" ]]; then
    echo "Error: Template file tidak ditemukan: $src"
    exit 1
  fi

  # Salin file dan ganti nama contoh ke nama feature baru
  sed \
    -e "s|source/feature/_example|source/feature/$FULL_PATH|g" \
    -e "s|/api/v1/_example|/api/v1/$FULL_PATH|g" \
    -e "s|Example|$FEATURE_CAMEL|g" \
    -e "s|EXAMPLE|$FEATURE_UPPER|g" \
    -e "s|example|$FEATURE_LOWER|g" \
    "$src" > "$dst"

  echo "  ✓ $(basename "$dst")"
}

# Proses setiap file dari template
copy_and_replace "$TEMPLATE_PATH/handler.go" "$FEATURE_PATH/handler.go"
copy_and_replace "$TEMPLATE_PATH/handler_impl.go" "$FEATURE_PATH/handler_impl.go"
copy_and_replace "$TEMPLATE_PATH/repository.go" "$FEATURE_PATH/repository.go"
copy_and_replace "$TEMPLATE_PATH/repository_impl.go" "$FEATURE_PATH/repository_impl.go"
copy_and_replace "$TEMPLATE_PATH/body/request.go" "$FEATURE_PATH/body/request.go"
copy_and_replace "$TEMPLATE_PATH/body/response.go" "$FEATURE_PATH/body/response.go"

echo ""
echo "Fitur berhasil dibuat!"
echo "Berikut langkah selanjutnya:"
echo "1. Sesuaikan dependensi di handler.go dan repository.go"
echo "2. Implementasikan logic di handler_impl.go dan repository_impl.go"
echo "3. Daftarkan route di source/service/route.go"
echo "4. Jalankan go build atau go test untuk memastikan tidak ada error"
echo ""
echo "Contoh pendaftaran route (di source/service/route.go):"
if [[ -n "$DOMAIN_PATH" ]]; then
  echo "  rg.Group(\"/api/v1/$FULL_PATH\").Handler(\"$FEATURE_UPPER\", handler.NewHandler(db))"
else
  echo "  rg.Group(\"/api/v1/$SCOPE/$FEATURE_LOWER\").Handler(\"$FEATURE_UPPER\", handler.NewHandler(db))"
fi
