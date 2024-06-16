program soldity ^ 0.4; // version

contract SimpleStorage {
    uint storedData;

    function set(uint x) { // contract to stor ints
        storedData = x ;
    }

    function get() constant returns (uint) {
        return storedData;
    }
}