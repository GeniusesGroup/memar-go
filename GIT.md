# GIT

## Rename many files
- Rename all `storage-services.go` to `storage.go`:    
```bash
find . -name 'storage-services.go' -exec sh -c 'git mv "$0" "${0%storage-services.go}storage.go"' {} \;

find . -name 'department*' -exec sh -c 'git mv -- "$0" "${0%department}org-department"' {} \;
```
    - Add `*` to end of `-name` to select start or end with.

## Move all directory items to another directory
```bash
git mv /path/new/old/* /path/new/.
```

## List only deleted files not commit yet
```bash
git diff --name-only --diff-filter D
```
