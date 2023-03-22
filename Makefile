RELEASE_VERSION=v$(shell cat package.json | jq --raw-output .version)
 
TARGET_BRANCH=master

publish:
	@npm publish

commit:
	@git commit -am "release ${RELEASE_VERSION}"

tag:
	@git tag ${RELEASE_VERSION}

push:
	@git push origin ${TARGET_BRANCH} ${RELEASE_VERSION}

release: publish commit tag push
	 