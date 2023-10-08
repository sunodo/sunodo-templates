package io.sunodo.app;

import io.cartesi.rollups.ApiException;
import io.cartesi.rollups.api.DefaultApi;
import io.cartesi.rollups.model.Advance;
import io.cartesi.rollups.model.Finish;
import io.cartesi.rollups.model.Inspect;
import io.cartesi.rollups.model.RollupRequest;
import io.cartesi.rollups.model.Finish.StatusEnum;
import io.cartesi.rollups.model.RollupRequest.RequestTypeEnum;

public class App {

    private static StatusEnum handleAdvance(Advance request) {
        System.out.println("Received advance request data " + request.getPayload());
        return StatusEnum.ACCEPT;
    }

    private static StatusEnum handleInspect(Inspect request) {
        System.out.println("Received inspect request data " + request.getPayload());
        return StatusEnum.ACCEPT;
    }

    public static void main(String[] args) {
        String rollupServer = System.getenv("ROLLUP_HTTP_SERVER_URL");
        System.out.println("HTTP rollup_server url is " + rollupServer);

        DefaultApi api = new DefaultApi();
        api.setCustomBaseUrl(rollupServer);

        Finish finish = new Finish().status(StatusEnum.ACCEPT);
        while (true) {
            try {
                RollupRequest request = api.finish(finish);
                RequestTypeEnum type = request.getRequestType();
                if (type == RequestTypeEnum.ADVANCE_STATE) {
                    finish = new Finish().status(handleAdvance(request.getData().getAdvance()));
                } else if (type == RequestTypeEnum.INSPECT_STATE) {
                    finish = new Finish().status(handleInspect(request.getData().getInspect()));
                } else {
                    System.out.println("Unknown request type: " + type);
                }

            } catch (ApiException e) {
                System.out.println(e);
            }
            try {
                Thread.sleep(5000);
            } catch (InterruptedException e) {
                System.out.println(e);
            }
        }
    }
}
