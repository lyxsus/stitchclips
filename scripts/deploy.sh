rsync -r --delete-after --quiet $TRAVIS_BUILD_DIR/stitchclips root@sadzeih.com:/src/stitchclips
ssh -tt root@sadzeih.com "systemctl restart stitchclips-api"