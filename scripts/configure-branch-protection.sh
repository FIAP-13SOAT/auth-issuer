#!/usr/bin/env bash
# =============================================================================
# Script: configure-branch-protection.sh
# Descrição: Configura branch protection no repositório auth-issuer via GitHub API
# Requisito: 10.1 — Exigir ao menos 1 aprovação de revisão na branch principal
#
# Variáveis de ambiente necessárias:
#   GITHUB_TOKEN  — Personal Access Token com permissão 'repo'
#   GITHUB_OWNER  — Owner/organização do repositório no GitHub
#
# Uso:
#   export GITHUB_TOKEN="ghp_..."
#   export GITHUB_OWNER="meu-usuario"
#   ./scripts/configure-branch-protection.sh
# =============================================================================

set -euo pipefail

REPO="auth-issuer"
BRANCH="master"
API_BASE="https://api.github.com"

# --- Validação de variáveis ---
if [ -z "${GITHUB_TOKEN:-}" ]; then
  echo "Erro: GITHUB_TOKEN não definido."
  echo "Exporte o token: export GITHUB_TOKEN=\"ghp_...\""
  exit 1
fi

if [ -z "${GITHUB_OWNER:-}" ]; then
  echo "Erro: GITHUB_OWNER não definido."
  echo "Exporte o owner: export GITHUB_OWNER=\"seu-usuario\""
  exit 1
fi

PROTECTION_URL="${API_BASE}/repos/${GITHUB_OWNER}/${REPO}/branches/${BRANCH}/protection"

echo "Configurando branch protection para ${GITHUB_OWNER}/${REPO} (branch: ${BRANCH})..."

# --- Aplicar branch protection ---
HTTP_STATUS=$(curl -s -o /tmp/bp-response.json -w "%{http_code}" \
  -X PUT \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  "${PROTECTION_URL}" \
  -d '{
    "required_status_checks": null,
    "enforce_admins": false,
    "required_pull_request_reviews": {
      "required_approving_review_count": 1
    },
    "restrictions": null
  }')

if [ "$HTTP_STATUS" -eq 200 ]; then
  echo "Branch protection configurada com sucesso."
  echo "  - Branch: ${BRANCH}"
  echo "  - Aprovações mínimas: 1"
else
  echo "Erro ao configurar branch protection (HTTP ${HTTP_STATUS}):"
  cat /tmp/bp-response.json
  exit 1
fi

# --- Verificação ---
echo ""
echo "Verificando configuração..."

VERIFY_STATUS=$(curl -s -o /tmp/bp-verify.json -w "%{http_code}" \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer ${GITHUB_TOKEN}" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  "${PROTECTION_URL}")

if [ "$VERIFY_STATUS" -eq 200 ]; then
  APPROVALS=$(cat /tmp/bp-verify.json | grep -o '"required_approving_review_count":[0-9]*' | head -1 | cut -d: -f2)
  echo "Verificação OK — Aprovações exigidas: ${APPROVALS:-N/A}"
else
  echo "Aviso: Não foi possível verificar (HTTP ${VERIFY_STATUS})."
fi

# --- Limpeza ---
rm -f /tmp/bp-response.json /tmp/bp-verify.json

echo ""
echo "Concluído."
