/* For license and copyright information please see the LEGAL file in the code repository */

package www

const returnHTMLPage = `
<!-- For license and copyright information please see the LEGAL file in the code repository -->
<!DOCTYPE html>
<html>

<head>
    <title>Your browser is outdated!</title>
</head>

<body>
    <style>
        main {
            position: absolute;
            top: 10%;
            left: 0;
            right: 0;
            bottom: 0;
            max-width: 100%;
            max-height: 100%;
            height: 200px;
            text-align: center;
            font-weight: 300;
            font-family: 'Roboto', Arial, sans-serif;
        }

        main a {
            text-decoration: none;
        }

        main p {
            font-size: 1.15rem;
            letter-spacing: -0.3px;
            line-height: 20px;
            font-weight: 500;
            text-transform: uppercase;
            color: #92989b;
        }

        .chip {
            display: inline-block;
            height: 64px;
            font-size: 20px;
            text-transform: uppercase;
            font-weight: 500;
            color: #616161;
            line-height: 64px;
            padding-right: 24px;
            border-radius: 32pt;
            background-color: #e4e4e4;
            margin-bottom: 5px;
            margin-right: 15px;
            letter-spacing: -0.5px;
            padding: 0 32px 0 12px;
            -webkit-box-shadow: 0 3px 3px 0 rgba(0, 0, 0, 0.14), 0 1px 7px 0 rgba(0, 0, 0, 0.12), 0 3px 1px -1px rgba(0, 0, 0, 0.2);
            box-shadow: 0 3px 3px 0 rgba(0, 0, 0, 0.14), 0 1px 7px 0 rgba(0, 0, 0, 0.12), 0 3px 1px -1px rgba(0, 0, 0, 0.2);
            -webkit-transition: all 0.1s ease-out;
            -moz-transition: all 0.1s ease-out;
            -o-transition: all 0.1s ease-out;
            transition: all 0.1s ease-out;
        }

        .chip:hover {
            color: #4CAF50;
            -webkit-box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.14), 0 1px 10px 0 rgba(0, 0, 0, 0.12), 0 2px 4px -1px rgba(0, 0, 0, 0.3);
            box-shadow: 0 4px 5px 0 rgba(0, 0, 0, 0.14), 0 1px 10px 0 rgba(0, 0, 0, 0.12), 0 2px 4px -1px rgba(0, 0, 0, 0.3);
        }

        .chip>img {
            float: left;
            margin: 0 12px 0 -12px;
            height: 64px;
            width: auto;
            border-radius: 50%;
        }
    </style>
    <main>
        <h2>Your browser is outdated!</h2>
        <div>
            <p>
                Don't take this personally, but your browser is too old to run <strong>the Platform</strong>.<br />
                We require HTML5 support to work.
            </p>
            <p>
                Unfortunately, your current browser doesn't support HTML5.<br />
                Please download or update a modern browser and come back soon!
            </p>
        </div>

        <div class="chips">
            <a href="https://www.google.com/chrome/browser/" title="Chrome browser download link">
                <div class="chip"><img src="https://www.google.com/chrome/static/images/chrome-logo.svg"
                        alt="Chrome browser logo">Chrome</div>
            </a>

            <a href="https://www.mozilla.org/en-US/firefox/new/" title="Mozilla browser download link">
                <div class="chip"><img
                        src="https://www.mozilla.org//media/img/logos/firefox/logo-quantum.9c5e96634f92.png"
                        alt="Mozilla browser logo">Firefox</div>
            </a>

            <a href="https://www.apple.com/safari/" title="Safari browser download link">
                <div class="chip"><img src="https://www.apple.com//v/safari/i/images/overview/safari_icon_large.png"
                        alt="Safari browser logo">Safari</div>
            </a>
        </div>
    </main>
</body>
`
