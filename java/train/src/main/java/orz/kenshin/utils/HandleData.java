package orz.kenshin.utils;

class Person1{
    public String name;
    public int age;
}

class Person2{
    public String name;
}

class App{
    public static void main(String[] args){
        Person1 p1 = new Person1();
        p1.age = 20;
        Person2 p2 = new Person2();
        p2.name = "zs";

        HandleData<Person1, Person2> handleData = new HandleData<Person1, Person2>();



        p1 =  handleData.add(p1, p2);
        System.out.println(p1.name);

    }
}


class HandleData<T1, T2>{
    public  <T1> add(T1 t1, T2 t2) {
        t1.name = t2.name;
        return t1;
        // return new Pair<K>(first, last);
    }
}