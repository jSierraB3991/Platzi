package org.jSierra3991.payments.service;

import org.jSierra3991.payments.models.PaymentRequest;
import org.jSierra3991.payments.models.PaymentResponse;

public interface PaymentGateWay {

    PaymentResponse requestPayment(PaymentRequest request);
}
