document.getElementById('emailForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    const subject = document.getElementById('subject').value;
    const recipient = document.getElementById('recipient').value;
    const body = document.getElementById('body').value;
    const resultBox = document.getElementById('result');

    try {
        const response = await fetch(`https://sendit-email-assistant.victoriousgrass-edd82384.westus2.azurecontainerapps.io/notifications/email`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "subject": subject,
                "recipient": recipient,
                "body": body
            })
        });

        if (response.status === 201) {
            resultBox.className = 'result-box success';
            resultBox.textContent = `Email sent successfully to ${recipient} with subject: "${subject}"!`;
        } else {
            const errorData = await response.json();
            resultBox.className = 'result-box error';
            resultBox.textContent = `Error sending email to ${recipient} with subject "${subject}": ${errorData.details || 'Failed to send email. Please try again.'}`;
        }
    } catch (error) {
        console.error('Error sending email:', error);
        resultBox.className = 'result-box error';
        resultBox.textContent = `Error sending email to ${recipient} with subject "${subject}": ${error.message || 'Error connecting to the server. Please try again.'}`;
    }
});

function clearForm() {
    document.getElementById('subject').value = '';
    document.getElementById('recipient').value = '';
    document.getElementById('body').value = '';
    document.getElementById('result').className = 'result-box';
    document.getElementById('result').textContent = '';
} 