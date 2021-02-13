# urler

Small tool to convert urls from stdin into json. Useful in combination with `jq`.
Example usage:

```
nomad% echo "https://github.com/rHermes/urler/pulls?q=is%3Apr+is%3Aopen+sort%3Aupdated-desc" | ./urler | jq
{
  "Scheme": "https",
  "Opaque": "",
  "User": null,
  "Host": "github.com",
  "Path": "/rHermes/urler/pulls",
  "RawPath": "",
  "ForceQuery": false,
  "RawQuery": "q=is%3Apr+is%3Aopen+sort%3Aupdated-desc",
  "Fragment": "",
  "RawFragment": "",
  "Values": {
    "q": [
      "is:pr is:open sort:updated-desc"
    ]
  },
  "Raw": "https://github.com/rHermes/urler/pulls?q=is%3Apr+is%3Aopen+sort%3Aupdated-desc"
}
```
