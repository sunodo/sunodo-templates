# TypeScript+SQLite DApp Template

This is a template for TypeScript Cartesi DApps. It uses node to execute the backend application.
The application entrypoint is the `src/index.ts` file. It is bundled with [esbuild](https://esbuild.github.io), but any bundler can be used.

This template is based on `typescript` template adding a SQlite database to it. the template already implements database creation and simple iterations to `add`, `delete` and `list` entries form a `product` table.

To add new entries is just required a Input with the following payload:

```json
{
  "id": "1",
  "name": "test",
  "action": "add"
}
```

To remove entries you just send a payload with action as `delete`:

```json
{
  "id": "1",
  "name": "test",
  "action": "delete"
}
```

To list entries you just need to call a inspect function on node machine:

```sh
 curl https://localhost:8080/inspect/products
```
