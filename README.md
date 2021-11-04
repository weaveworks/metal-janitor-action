# Metal Janitor Action

A GitHub Action to cleanup Equinix Metal projects.

## Inputs

### `metal-auth-token`

**Required** The Equinix Metal API token to use. See the [docs](https://metal.equinix.com/developers/docs/accounts/users/#api-keys) for further information on API keys.

### `project-names`

**Required** The names of the Equinix Metal projects to cleanup.

You can specify multiple projects to cleanup by comma separating them. Projects will be cleanup up in the order specified: For example: `projecta,projectb`.

You can delete all projects by using the value **DELETEALL**. _Note: this will delete
all projects in the Org, or all projects that your API key can access._

## Example Usage

```yaml
jobs:
  cleanup:
    runs-on: ubuntu-latest
    name: Cleanup projects
    steps:
      - name: Cleanup projects
        uses: weaveworks/metal-janitor-action@b0373b3a6a8bb1e6573616040121c22ba230cdc8
        with:
          metal-auth-token: ${{ secrets.METAL_AUTH_TOKEN }}
          project-names: 'project1,project2'
```
