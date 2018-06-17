// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num);

class Solution {
public:
    int guessNumber(int n) {
        int low = 1;
        int higt = n;
        int mid = low + (higt-low)/2;

        while(low <= higt){
           int g = guess(mid);
           if(g == 0) {
               break;
           } else if(g>0){
               low = mid + 1;
               
           } else {
               higt = mid - 1;
           }

           mid = low + (higt-low)/2;
        }
        
        return mid;
    }
};
