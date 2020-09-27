local_resource('protos', 'earth +protos', deps=[
    'pkg/game/rpcrepository/repo.proto',
    'pkg/turn/rpcturn/current.proto'
])

local_resource(
    'binary-build', 
    'earth +binaries', 
    deps=[
        'cmd', 'internal', 'pkg'
    ], 
    resource_deps=['protos']
)

custom_build(
    'api-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/api/+docker',
    ['./.output/api'],
    live_update = [
        sync('./.output/api', '/root/app'),
        run('./restart.sh'),
    ]
)
k8s_yaml('deploy/api/deploy.yaml')
k8s_resource('api', port_forwards="8081:8080", resource_deps=['binary-build'])

custom_build(
    'web-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/web/+docker',
    ['./.output/web'],
    live_update = [
        sync('./.output/web', '/root/app'),
        run('./restart.sh'),
    ]
)
k8s_yaml('deploy/web/deploy.yaml')
k8s_resource('web', port_forwards="8080:8080", resource_deps=['binary-build'])

custom_build(
    'gamerepo-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/gamerepo/+docker',
    ['./.output/gamerepo'],
    live_update = [
        sync('./.output/gamerepo', '/root/app'),
        run('./restart.sh'),
    ]
)
k8s_yaml('deploy/gamerepo/deploy.yaml')
k8s_resource('gamerepo', port_forwards=["8082:8080", "8083:8081"], resource_deps=['binary-build'])

custom_build(
    'currentturn-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/currentturn/+docker',
    ['./.output/currentturn'],
    live_update = [
        sync('./.output/currentturn', '/root/app'),
        run('./restart.sh'),
    ]
)
k8s_yaml('deploy/currentturn/deploy.yaml')
k8s_resource('currentturn', port_forwards=["8084:8080", "8085:8081"], resource_deps=['binary-build'])

custom_build(
    'grid-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/grid/+docker',
    ['./.output/grid'],
    live_update = [
        sync('./.output/grid', '/root/app'),
        run('./restart.sh'),
    ]
)
k8s_yaml('deploy/grid/deploy.yaml')
k8s_resource('grid', port_forwards=["8086:8080", "8087:8081"], resource_deps=['binary-build'])