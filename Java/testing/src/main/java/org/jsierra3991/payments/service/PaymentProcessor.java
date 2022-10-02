package org.jsierra3991.payments.service;

import lombok.RequiredArgsConstructor;
import org.jsierra3991.payments.models.PaymentRequest;
import org.jsierra3991.payments.models.PaymentStatus;

@RequiredArgsConstructor
public class PaymentProcessor {

    private final PaymentGateWay paymentGateWay;

    public boolean makePayment(double amount) {
        var paymentStatus = paymentGateWay.requestPayment(PaymentRequest.builder().amount(amount).build()).getStatus();
        return paymentStatus.equals(PaymentStatus.OK);
    }
}
