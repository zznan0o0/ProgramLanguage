class GenericTest{
    public static void main(String args[]){
        // String str = genericTest();
        // System.out.println(str);
        GenricFun<String, Integer> genricFun = new GenricFun<String, Integer>();
        System.out.println(genricFun.toString(123456));
    }

    // public static <S> S genericTest(){
    //     return "AABB";
    //     // return (S) Integer.toString(arg1);
    // }
}

class GenricFun<S,I>{
    public S toString(I i){
        // return (S) String.valueOf(i);
        return (S) (i+"");
    }
}