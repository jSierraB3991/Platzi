package org.jSierra3991.service;

import org.jSierra3991.payments.models.PaymentResponse;
import org.jSierra3991.payments.service.PaymentGateWay;
import org.jSierra3991.payments.service.PaymentProcessor;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;

import static org.jSierra3991.payments.models.PaymentStatus.ERROR;
import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.when;
import static org.jSierra3991.payments.models.PaymentStatus.OK;

//org.mockito.junit.jupiter.MockitoExtension;
//org.junit.jupiter.api.extension.ExtendWith;
//jupiter @ExtendWith(MockitoExtension.class)
@RunWith(MockitoJUnitRunner.class)
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

    @Test
    public void payment_is_worng(){
        when(paymentGateWay.requestPayment(any()))
                .thenReturn(PaymentResponse.builder().status(ERROR).build());
        boolean result = paymentProcessor.makePayment(1000);
        assertFalse(result);
    }
}