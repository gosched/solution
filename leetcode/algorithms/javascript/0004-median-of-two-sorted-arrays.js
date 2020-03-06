/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
    let length = nums1.length + nums2.length;
    let mid = Math.floor(length / 2);
    let midLeft = mid - 1;
    let i = 0, j = 0, k = 0;
    let last = 0, left = 0, right = 0;

    while (k <= mid && i < nums1.length && j < nums2.length) {
        if (nums1[i] < nums2[j]) {
            last = nums1[i];
            i++
        } else {
            last = nums2[j];
            j++;
        }
        if (k == midLeft) {
            left = last;
        }
        if (k == mid) {
            right = last;
        }
        k++;
    }

    while (k <= mid && i < nums1.length) {
        last = nums1[i];
        i++
        if (k == midLeft) {
            left = last;
        }
        if (k == mid) {
            right = last;
        }
        k++;
    }

    while (k <= mid && j < nums2.length) {
        last = nums2[j];
        j++;
        if (k == midLeft) {
            left = last;
        }
        if (k == mid) {
            right = last;
        }
        k++;
    }

    if (length % 2 == 0) {
        return (left + right) / 2;
    }
    return right;
};