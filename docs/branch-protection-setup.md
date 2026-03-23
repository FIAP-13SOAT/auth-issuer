# Configuração de Branch Protection — auth-issuer

## Branch Principal

- **Branch:** `master`

## Requisitos de Proteção

- Exigir ao menos **1 aprovação de revisão** (pull request review) antes do merge na branch `master`.

---

## Opção 1: Configuração via GitHub UI

1. Acesse o repositório `auth-issuer` no GitHub.
2. Vá em **Settings** → **Branches**.
3. Em **Branch protection rules**, clique em **Add rule** (ou **Add branch ruleset**).
4. No campo **Branch name pattern**, insira: `master`
5. Marque a opção **Require a pull request before merging**.
6. Em **Required approvals**, defina o valor mínimo como **1**.
7. (Opcional) Marque **Require status checks to pass before merging** e adicione os checks do pipeline CI/CD.
8. Clique em **Create** ou **Save changes**.

---

## Opção 2: Configuração via GitHub API (curl)

Execute o comando abaixo substituindo `OWNER` pelo nome do owner/organização e `GITHUB_TOKEN` por um token com permissão `repo` (ou `admin:org` para organizações):

```bash
curl -X PUT \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/OWNER/auth-issuer/branches/master/protection \
  -d '{
    "required_status_checks": null,
    "enforce_admins": false,
    "required_pull_request_reviews": {
      "required_approving_review_count": 1
    },
    "restrictions": null
  }'
```

### Verificação

Para verificar se a proteção foi aplicada:

```bash
curl -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/OWNER/auth-issuer/branches/master/protection
```

---

## Opção 3: Script automatizado

Use o script `scripts/configure-branch-protection.sh` disponível neste repositório:

```bash
export GITHUB_TOKEN="ghp_seu_token_aqui"
export GITHUB_OWNER="seu_usuario_ou_org"
./scripts/configure-branch-protection.sh
```

---

## Referências

- [GitHub Docs — Managing a branch protection rule](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/managing-a-branch-protection-rule)
- [GitHub REST API — Branch Protection](https://docs.github.com/en/rest/branches/branch-protection)
