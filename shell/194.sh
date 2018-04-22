awk '
{
    for(i=1;i<=NF;i++) {
        if(NR==1){
            arr[i] = $i
        } else {
            arr[i] = arr[i] " "  $i
        }
    }
}
END {
   for(itam in arr)
      print(arr[itam])
    
}
' file.txt
