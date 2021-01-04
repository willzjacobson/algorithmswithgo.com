// NOTE: Seems to be broken. This counting sort implementation seems not to be stable
const objects = [{ date: '2-20-18', val: 4, }, { date: '4-20-17', val: 6, }, { date: '3-20-17', val: 5, }, { date: '6-2-17', val: 7, }, { date: '2-18-18', val: 3, }, { date: '1-20-17', val: 2, }, { date: '6-3-15', val: 1, }];
radixSort(objects);
console.log('Objects using Radix Sort:', objects);

function radixSort(arr) {
  for (let i = 0; i < arr.length; i++) {
    const [ month, day, year ] = arr[i].date.split('-');
    arr[i].month = Number(month);
    arr[i].day = Number(day);
    arr[i].year = Number(year);
  }

  countingSort(arr, 31, 'day');
  console.log('post day:', arr)
  countingSort(arr, 12, 'month');
  console.log('post month:', arr)
  countingSort(arr, 18, 'year');

  for (let i = 0; i < arr.length; i++) {
    delete arr[i].day;
    delete arr[i].month;
    delete arr[i].year;
  }
}

// We are guaranteed that each element of 'nums' is an integer x where 0 <= x <= k.
function countingSort(objects, k, field) {
  const out = Array(objects.length);
  const store = Array(k+1).fill(0);
  // make each idx of 'store' hold the # of elements in 'nums' that equal that idx
  for (const o of objects) {
    store[o[field]]++;
  }
  // make each idx of 'store' hold the # of elements that come before it
  for (let i = 1; i < store.length; i++) {
    store[i] += store[i-1];
  }
  // place elements in correct positions in 'out' array
  for (let i = 0; i < objects.length; i++) {
    const o = objects[i];
    out[store[o[field]]-1] = o;
    store[o[field]]--;
  }
  // overwrite input array using 'out' array, which is sorted
  for (let i = 0; i < out.length; i++) {
    objects[i] = out[i];
  }
}

console.log('\n=== counting sort test ===')
const arr1 = [{ val: 4, }, { val: 5, }, { val: 6, }, { val: 3, }, { val: 2, }, { val: 1, }];
countingSort(arr1, 7, 'val');
console.log('arr1:', arr1);
