package org.example.api;

import java.io.Serializable;

/**
 * @author: liushitao
 * @description:
 * @date: 2025/3/17 15:48
 * @version: 1.0
 */
public class RpcResult<T> implements Serializable {
    private int code;
    private String message;
    private T data;


    public RpcResult(int code, String message, T data) {
        this.code = code;
        this.message = message;
        this.data = data;
    }

    public static <T> RpcResult<T> success(T data) {
        return new RpcResult<>(0, "success", data);
    }

}
