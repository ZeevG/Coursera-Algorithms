
function merge(left, right){

    var sorted = [];

    while (left.length !== 0 && right.length !== 0){

        if (left[0] < right[0]){
            sorted.push(left.shift());
        } else {
            sorted.push(right.shift());
        }
    }

    while (right.length !== 0){
        sorted.push(right.shift());
    }

    while (left.length !== 0){
        sorted.push(left.shift());
    }

    return sorted;

}

function mergeSort(unsorted){

    console.log("Unsorted ", unsorted);


    if (unsorted.length === 1){
        return unsorted;
    }

    var left = unsorted.slice(0, Math.floor(unsorted.length/2));
    var right = unsorted.slice(Math.floor(unsorted.length/2), unsorted.length);

    left = mergeSort(left);
    right = mergeSort(right);

    sorted = merge(left, right);

    return sorted;

}

unsorted = [
    34,
    23,
    423,
    356,
    43,
    4,
    2,
    32,
    41,
    2154,
    5,
    63,
    43,
    123,
    42,
    34,
    52,
    35,
    345,
];

console.log(unsorted);
sorted = mergeSort(unsorted);
console.log(sorted);

