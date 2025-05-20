const apiUrl = import.meta.env.APLICATION_URL;

if (!apiUrl) {
    console.error('API URL is not configured. Please check your environment variables.');
    document.getElementById('result').className = 'result-box error';
    document.getElementById('result').textContent = 'Application configuration error. Please contact support.';
}

document.getElementById('emailForm').addEventListener('submit', async function(e) {
    e.preventDefault();

    if (!apiUrl) {
        return;
    }

    const subject = document.getElementById('subject').value;
    const recipient = document.getElementById('recipient').value;
    const body = document.getElementById('body').value;
    const resultBox = document.getElementById('result');

    try {
        const response = await fetch(`${apiUrl}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                subject,
                recipient,
                body
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