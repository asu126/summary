class RecentCounter {
    private List<Integer> list; 
    public RecentCounter() {
        list = new ArrayList<Integer>();
    }
    
    public int ping(int t) {
        list.add(t);
        int result = 0;
        
        // for(Integer i:list){
        //     if(t-i<=3000){
        //         result++;
        //     }
        // }
        for(int i=list.size()-1; i>=0; i--) { 
            Integer a = list.get(i);
            if(t-a<=3000){
                result++;
            }else{
                break;
            }
        }
        
        return result;
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */