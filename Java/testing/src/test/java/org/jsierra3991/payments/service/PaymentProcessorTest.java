package org.jsierra3991.payments.service;


import org.jsierra3991.payments.models.PaymentRequest;
import org.jsierra3991.payments.models.PaymentResponse;
import org.junit.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;

import static org.junit.Assert.assertTrue;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.jsierra3991.payments.models.PaymentStatus.OK;

public class PaymentProcessorTest {

    @InjectMocks
    private PaymentProcessor paymentProcessor;
    @Mock
    private PaymentGateWay paymentGateWay;

    @Test
    public void payment_is_correct(){
        when(paymentGateWay.requestPayment(any()))
                .thenReturn(PaymentResponse.builder().status(OK).build());
        boolean result = paymentProcessor.makePayment(1000);
        assertTrue(result);
    }
}