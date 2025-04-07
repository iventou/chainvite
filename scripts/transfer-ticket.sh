chainvited tx ticket transfer-ticket \
  --ticket-id "[ticket-id]" \
  --new-owner "$(chainvited keys show bob -a)" \
  --from alice \
  --chain-id chainvite \
  --gas auto \
  --gas-adjustment 1.3 \
  --yes
