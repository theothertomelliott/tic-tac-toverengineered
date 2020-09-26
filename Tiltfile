local_resource('protos', 'earth +protos', deps=[
    'pkg/game/rpcrepository/repo.proto',
    'pkg/turn/rpcturn/current.proto'
])

custom_build(
    'api-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/api/+docker',
    ['cmd', 'internal', 'pkg'])
k8s_yaml('deploy/api/deploy.yaml')
k8s_resource('api', port_forwards="8081:8080")

custom_build(
    'web-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/web/+docker',
    ['cmd', 'internal', 'pkg'])
k8s_yaml('deploy/web/deploy.yaml')
k8s_resource('web', port_forwards="8080:8080")

custom_build(
    'gamerepo-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/gamerepo/+docker',
    ['cmd', 'internal', 'pkg'])
k8s_yaml('deploy/gamerepo/deploy.yaml')
k8s_resource('gamerepo', port_forwards=["8082:8080", "8083:8081"])

custom_build(
    'currentturn-image',
    'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/currentturn/+docker',
    ['cmd', 'internal', 'pkg'])
k8s_yaml('deploy/currentturn/deploy.yaml')
k8s_resource('currentturn', port_forwards=["8084:8080", "8085:8081"])