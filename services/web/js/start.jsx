import React from "react";
import ReactDOM from "react-dom/client";
var TicTacToe = require('tic_tac_toe');

var api = new TicTacToe.DefaultApi(new TicTacToe.ApiClient("http://localhost:8081"));

class Grid extends React.Component {
    constructor(props) {
        super(props);
        this.state = {grid: false};
        var g = this;
        api.gameGrid(props.id,
            function (error, data, response) {
                if (error) {
                    console.error(error);
                } else {
                    console.log(data.grid);
                    g.setState({grid: data.grid});
                }
            });
    }

    render() {
        return <div className="grid grid-cols-3 gap-2 my-5">
                <GridInner grid={this.state.grid}/>
            </div>;
    }
}

function GridInner(props) {
    if (props.grid) {
        var spaces = [];
        props.grid.forEach((row, i) => {
            row.forEach((space, j) => {
                spaces.push(
                    <GridSpace
                        key={"space:" + i + "," + j}
                        content={space}
                        play={() => {
                            doPlay(i,j);
                        }}
                    />);
            });
        });
        return <>{spaces}</>;
    } else {
        return <p>No grid</p>;
    }
}

function GridSpace(props) {
    var content = <>&nbsp;&nbsp;</>;
    if (props.content != "") {
        content = props.content;
    }
    return <div className="block">
        <button
            className="block w-full rounded-md text-white font-extrabold text-center hover:bg-blue-600 bg-blue-500 p-7 hover:shadow-lg focus:outline-none"
            type="button"
            onClick={props.play}
        >
            {content}
        </button>
    </div>;
}

var grid = document.getElementById("grid");
if (grid) {
    var gid = grid.dataset.gameid;
    ReactDOM.createRoot(grid).render(<Grid id={gid}/>);
}

function doPlay(x,y) {
    var gid = grid.dataset.gameid
    api.currentPlayer(gid, function (error, data, response) {
        if (error) {
            console.error(error);
        } else {
            play(gid,data,x,y);
        }
    })
}

function play(gid,p,x,y){
    // TODO: Handle finished
    location.href = "/" + gid + "/play?player=" + p + "&pos={\"i\":" + x + ",\"j\":" + y + "}";
}