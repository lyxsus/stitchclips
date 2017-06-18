rsync -r --delete-after $TRAVIS_BUILD_DIR/stitchclips root@sadzeih.com:/srv/stitchclips
ssh -tt root@sadzeih.com "systemctl restart stitchclips-api"
exit 0