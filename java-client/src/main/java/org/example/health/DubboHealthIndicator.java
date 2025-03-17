package org.example.health;

import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.stereotype.Component;

@Component
public class DubboHealthIndicator implements HealthIndicator {

    @Override
    public Health health() {
        try {
            // 这里可以添加自定义的健康检查逻辑
            return Health.up()
                    .withDetail("dubboStatus", "running")
                    .withDetail("serviceCount", getServiceCount())
                    .build();
        } catch (Exception e) {
            return Health.down()
                    .withDetail("error", e.getMessage())
                    .build();
        }
    }

    private int getServiceCount() {
        // 这里可以实现获取当前发布的服务数量的逻辑
        return 1; // 示例返回值
    }
} 