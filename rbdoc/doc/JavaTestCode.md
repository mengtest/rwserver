###tcp JAVA 测试代码 ,服务端对半包粘包已处理
1.启动rbgate的GateServer  
2.运行测试代码
```
import java.io.*;
import java.net.Socket;
import java.net.UnknownHostException;


public class TestMain {
    public static void main(String[] args) {
        try {
            //创建一个客户端socket
            Socket socket = new Socket("127.0.0.1",1024);
            //向服务器端传递信息
            DataOutputStream ots = new DataOutputStream(socket.getOutputStream());
            String content="{23213891271237192731928371231233012931231wewqeqweqweqwedfsdfsdfsdfqeweqweqeqwe209381293891289381902389}";
            content="{\"cmd\":\"1000\"}";
            for (int i = 0; i < 10; i++) {
                //java默认是通信进行大端传输，所以这里不用管大小端
                byte[] wb=content.getBytes("UTF-8");
                int ln=wb.length;
                byte[] lengthbytes = integerToBytes(ln, 2);
                ots.write(unitByteArray(lengthbytes,wb));
                ots.flush();
            }
            ots.close();
            socket.close();
        } catch (UnknownHostException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static byte[] integerToBytes(int num, int len) {
        byte[] bytes = ByteBuffer.allocate(len).putInt(num).array();
        return  bytes;
    }

    public static byte[] unitByteArray(byte[] byte1,byte[] byte2){
        byte[] unitByte = new byte[byte1.length + byte2.length];
        System.arraycopy(byte1, 0, unitByte, 0, byte1.length);
        System.arraycopy(byte2, 0, unitByte, byte1.length, byte2.length);
        return unitByte;
    }

}



```