package org.jsierra3991.payments.service;

import org.jsierra3991.payments.models.PaymentRequest;
import org.jsierra3991.payments.models.PaymentResponse;

public interface PaymentGateWay {

    PaymentResponse requestPayment(PaymentRequest request);
}
