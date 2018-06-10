// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int left =1;
        int rigth =n;
        int mid;
        
        while(left<=rigth){
            mid = left + (rigth-left)/2;

            if(isBadVersion(mid)){
                if(mid-1 >= 1 && !isBadVersion(mid-1)){
                    return mid;
                }else{
                   rigth = mid-1;
                }
            }else{
                 left = mid+1;
            }
        }
        return mid;
    }
};
