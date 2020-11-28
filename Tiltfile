update_settings(max_parallel_updates=1)

local_resource(
    'nginx',
    cmd='kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.41.2/deploy/static/provider/cloud/deploy.yaml'
)

local_resource(
    'tests',
    cmd='go test --short ./...',
    deps=['.'],
    ignore=['.output']
)

local_resource(
    'tests (long)',
    cmd='go test ./...',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
)

local_resource(
    'play',
    cmd='go run ./bot/cmd/onegame http://localhost:8081',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
)

local_resource(
    'protos', 
    'earth +protos', 
    deps=[
        'pkg/game/rpcrepository/repo.proto',
        'pkg/game/rpcrepository/Earthfile',
        'pkg/turn/rpcturn/current.proto',
        'pkg/turn/rpcturn/controller.proto',
        'pkg/turn/rpcturn/Earthfile',
        'pkg/grid/rpcgrid/grid.proto',
        'pkg/grid/rpcgrid/Earthfile',
        'pkg/space/rpcspace/space.proto',
        'pkg/space/rpcspace/Earthfile'
    ],
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
)

secrets = read_yaml("secrets.yaml")

# Load the base Helm chart for all resources
k8s_yaml(helm(
    'charts/tic-tac-toe',
    namespace='tictactoe',
    set=[
        "honeycomb.api_key=" + secrets["honeycomb"]["api_key"], 
        "honeycomb.dataset=tictactoe-dev",
        "mongodb.statefulset=true",
        ],
))

def server(name, port_forwards=[]):
    local_resource(
        name+"-build",
        'earth ./' + name + '/build/+build',
        #'GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/theothertomelliott/tic-tac-toverengineered/common/version.Version=tilt" -o ./.output/' + name + '/app ./' + name + '/cmd/' + name,
        deps = [name, "common"],
    )
    custom_build(
        "docker.io/tictactoverengineered/"+name,
        'earth --build-arg IMAGE_REF=$EXPECTED_REF ./' + name + '/build+docker',
        ['./.output/'+name],
        live_update = [
            sync('./.output/'+name, '/root/app'),
            run('./restart.sh'),
        ]
    )
    k8s_resource(name, port_forwards=port_forwards, resource_deps=[name+'-build'])

server("api", "8081:8080")
server("web", "8080:8080")
server("gamerepo", ["8082:8080", "8083:8081"])
server("currentturn", ["8084:8080", "8085:8081"])
server("grid",["8086:8080", "8087:8081"])
server("checker",["8088:8080", "8089:8081"])
server("turncontroller",["8090:8080", "8091:8081"])
server("bot")
server("space")

k8s_resource("mongodb-standalone", port_forwards="27017:27017")