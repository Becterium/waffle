.PHONY: api
# generate api
api:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# generate wire
wire:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'


.PHONY: image-conf
# move image conf to /data
image-conf:
	cp -r /mnt/hgfs/DevopsShareDir/waffle/deploy/images-conf/. /data/