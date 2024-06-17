// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

/// @title Greeter Contract
/// @custom:dev-run-script hello_world.js
contract Greeter {
    // variable of type string
    string public yourName;

    constructor() {
        /**
        set "Hello" in variable
        **/
        yourName = 'Hello';
    }

    function set(string memory name) public {
        /**
        add name to yourName
        **/
        yourName = name;
    }

    function hello() public view returns(string memory) {
        /**
        return full string
        **/
        return yourName;
    }
}
