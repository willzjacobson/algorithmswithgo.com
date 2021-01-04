const nums1 = [3, 3, 2, 1, 3, 5, 4, 1];
const nums2 = [6, 5, 4, 3, 3, 5, 4, 1];

function countingSort(nums, k) {
  const store = Array(k+1).fill(0);
  const out = Array(nums.length);

  for (const n of nums) {
    store[n]++;
  }
  for (let i = 1; i < store.length; i++) {
    store[i] += store[i-1];
  }

  for (const n of nums) {
    out[store[n]-1] = n;
    store[n]--;
  }

  return out;
}

console.log('1:', countingSort(nums1, 5));
console.log('2:', countingSort(nums2, 6));
