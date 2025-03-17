/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.example.service.impl;


import org.apache.dubbo.config.annotation.DubboService;
import org.example.api.GreetRequest;
import org.example.api.GreetResponse;
import org.example.api.GreetingsService;
import org.example.api.RpcResult;

@DubboService(
        version = "1.0.0",
        group = "dubbo",
        timeout = 3000,
        retries = 0
)
public class GreetingsServiceImpl implements GreetingsService {
    @Override
    public GreetResponse greet(GreetRequest req) {
        GreetResponse response = new GreetResponse();
        response.setGreeting("hi, " + req.getName());
        return response;
    }

    @Override
    public String sayHi(String name) {
        return name;
    }

    @Override
    public String sayHi2(GreetRequest req) {
        return req.getName();
    }

    @Override
    public GreetResponse sayHi3(String name) {
        GreetResponse response = new GreetResponse();
        response.setGreeting("hi, " + name);
        return response;
    }

    @Override
    public RpcResult<GreetResponse> sayHiGeneric(GreetRequest req) {
        GreetResponse response = new GreetResponse();
        response.setGreeting("hi, " + req.getName());
        return RpcResult.success(response);
    }
}
