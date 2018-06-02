#include <iostream>
#include <stdint.h>

// class Solution {
// public:
// 	int hammingWeight(uint32_t n) {
// 	    int result = 0;
// 	    uint32_t temp = 1;
	    
// 	    for(int i = 0; i < 32; i++){
// 	        if(n&temp != 0){
// 	        	// std::cout << temp << "  " << (n&temp) << " "<< (n&temp != uint32_t(0)) << std::endl;
// 	            result++;
// 	        }
	        
// 	        n = n>>1;
// 	    }
	        
// 	    return result;
// 	}
// };

int hammingWeight(uint32_t n) {
    int result = 0;
    uint32_t temp = 1;
    
    for(int i = 0; i < 32; i++){
        if(n&temp != 0){
        	// std::cout << temp << "  " << (n&temp) << " "<< (n&temp != uint32_t(0)) << std::endl;
            result++;
        }
        
        n = n>>1;
    }
        
    return result;
}

int main(int argc,char *argv[])
{
    std::cout << hammingWeight(2) << std::endl;
    return(0);
}