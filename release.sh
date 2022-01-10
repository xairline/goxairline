git pull --tags --force
git tag -l
read -p "new version(ex: 0.2.0):" version
git tag v${version} -f
git-chglog -o CHANGELOG.md
read -p " REVIEW CHANGELOG"
git add CHANGELOG.md
git commit -m "ci(release): bump up version: $version"
git push
git push origin --tags -f