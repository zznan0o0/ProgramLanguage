import java.lang.reflect.Method;

class GenericTest {
    public static void main(String args[]) {
        // String str = genericTest();
        // System.out.println(str);
        GenricFun<String, Integer> genricFun = new GenricFun<String, Integer>();
        System.out.println(genricFun.toString(123456));

        Animai cat = new Animai("猫", 5);
        // String name = ;
        System.out.println(HandleData.getFieldValueByName("name", cat));
    }

    // public static <S> S genericTest(){
    // return "AABB";
    // // return (S) Integer.toString(arg1);
    // }
}

class GenricFun<S, I> {
    public S toString(I i) {
        // return (S) String.valueOf(i);
        return (S) (i + "");
    }
}

class Animai{
    public String name;
    public int age;
    public Animai(String name, int age){
        this.name = name;
        this.age = age;
    }
}

class HandleData {
    public static Object getFieldValueByName(String fieldName, Object o) {
        try {
            String firstLetter = fieldName.substring(0, 1).toUpperCase();
            String getter = "get" + firstLetter + fieldName.substring(1);
            Method method = o.getClass().getMethod(getter, new Class[] {});
            Object value = method.invoke(o, new Object[] {});
            return value;
        } catch (Exception e) {
            System.out.println("获取属性值失败！" + e);
            // logger.error("获取属性值失败！" + e, e);
        }
        return null;
    }
}