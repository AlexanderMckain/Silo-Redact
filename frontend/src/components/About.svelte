<script>
  /**
   * Central About hub: mission (existing sections) + legal archive (#eula, #privacy, #terms).
   * Use with <Footer /> fixed at bottom; this root includes bottom padding so content clears the footer.
   * Legal text is for product disclosure — have counsel review before production reliance.
   */
  export let appVersion = ''

  const githubUrl = 'https://github.com/AlexanderMckain/Silo-Redact'
  const contactEmail = 'support@siloredact.com'

  /** Shared Tailwind classes: readable legal body, sans headings, mono for “system” callouts */
  const legalProse =
    'prose prose-invert prose-silo max-w-none prose-sm sm:prose-base ' +
    'prose-headings:font-sans prose-headings:font-semibold prose-headings:tracking-tight ' +
    'prose-p:font-sans prose-li:font-sans prose-strong:text-zinc-100 ' +
    'prose-a:text-cyan-400 prose-a:no-underline hover:prose-a:underline'
</script>

<div
  class="about-root text-left text-zinc-200 antialiased selection:bg-cyan-500/30 selection:text-white"
  aria-label="About SiloRedact"
>
  <!-- Mission & product narrative (existing sections) -->
  <header class="mb-8 border-b border-silo-border pb-6" id="top">
    <p class="mb-1 font-mono text-xs font-semibold uppercase tracking-widest text-silo-accent">
      Mission
    </p>
    <h1 class="font-sans text-xl font-semibold leading-snug text-white sm:text-2xl">
      Stop pasting sensitive data:<span class="text-silo-muted"> </span>Secure Your Data in the Age of AI.
    </h1>
  </header>

  <section class="mb-8 space-y-3 font-sans text-sm leading-relaxed text-zinc-300">
    <p>
      <span class="font-medium text-zinc-100">The problem.</span>
      Pasting logs, PII, secrets, or credentials into chatbots (LLMs), cloud support portals, or ticketing
      systems often copies that material into
      <strong class="text-zinc-100">remote data centers</strong>
      you do not control—expanding your attack surface with every paste.
    </p>
    <p>
      <span class="font-medium text-zinc-100">The threat.</span>
      Beyond routine data handling, attackers use
      <strong class="text-zinc-100">prompt injection</strong>
      and related techniques to manipulate AI into
      <strong class="text-zinc-100">exfiltrating or leaking</strong>
      context the model has seen—including sensitive text users pasted earlier—raising the stakes for
      <strong class="text-zinc-100">data breaches</strong>
      and unintended disclosure.
    </p>
    <p>
      <span class="font-medium text-zinc-100">The solution.</span>
      SiloRedact delivers
      <strong class="text-silo-accent">pre-flight, Local-First Architecture</strong>:
      <strong class="text-zinc-100"
        >all file redaction and processing occur 100% on your local machine. Your data never leaves your
        device</strong
      >
      for that workload—content is scrubbed before you paste, neutralizing much of that risk at the moment of copy/paste.
    </p>
    <p class="text-xs text-silo-muted">
      Built for
      <strong class="text-zinc-400">DevOps</strong>,
      <strong class="text-zinc-400">cybersecurity</strong>, and
      <strong class="text-zinc-400">IT</strong>
      professionals who need to share diagnostics without oversharing.
    </p>
  </section>

  <section
    class="mb-8 rounded-lg border border-cyan-900/50 bg-cyan-950/30 px-4 py-3 font-mono text-xs leading-relaxed text-cyan-100/90"
    aria-label="Security model summary"
  >
    <h2 class="mb-2 font-sans text-xs font-semibold uppercase tracking-wide text-cyan-400">
      Technical disclosure — security model
    </h2>
    <ul class="list-inside list-disc space-y-1.5 text-cyan-100/80">
      <li>
        <strong class="text-cyan-200/90">Local-First Architecture</strong> — file redaction and processing
        stay on-device; plaintext is not sent to SiloRedact for cloud scrubbing.
      </li>
      <li>Designed as a pre-paste control: you choose what leaves the machine after redaction.</li>
      <li>Stack: Go backend + Wails + Svelte UI — on-device trust boundary, desktop performance.</li>
    </ul>
    <p class="mt-3 border-t border-cyan-900/40 pt-3 text-cyan-100/70">
      <strong class="text-cyan-300">Note:</strong>
      SiloRedact is open source. Core redaction does not require network access. If you install from the Microsoft
      Store, Windows may connect to Microsoft servers for app updates — a system-level process that does not involve
      your processed data.
    </p>
  </section>

  <section class="mb-10" aria-labelledby="pillars-heading">
    <h2
      id="pillars-heading"
      class="mb-3 font-mono text-xs font-semibold uppercase tracking-widest text-silo-muted"
    >
      Three pillars
    </h2>
    <div class="grid gap-3 sm:grid-cols-3">
      <article class="rounded-md border border-silo-border bg-zinc-900/50 p-3 font-sans">
        <h3 class="mb-1 text-sm font-semibold text-white">Local-First Architecture</h3>
        <p class="text-xs leading-relaxed text-zinc-400">
          All file redaction and processing occur 100% on your machine. Your data never leaves your device for
          that workload—no cloud round-trip for scrubbing; telemetry on your content isn’t the product model.
        </p>
      </article>
      <article class="rounded-md border border-silo-border bg-zinc-900/50 p-3 font-sans">
        <h3 class="mb-1 text-sm font-semibold text-white">Privacy-aware</h3>
        <p class="text-xs leading-relaxed text-zinc-400">
          Purpose-built for <strong class="text-zinc-300">pre-flight</strong> scrubbing before pasting into LLMs
          (ChatGPT, Claude) or ticketing (Jira, Zendesk).
        </p>
      </article>
      <article class="rounded-md border border-silo-border bg-zinc-900/50 p-3 font-sans">
        <h3 class="mb-1 text-sm font-semibold text-white">Zero-trust mindset</h3>
        <p class="text-xs leading-relaxed text-zinc-400">
          Assume every paste destination is hostile.
          <strong class="text-zinc-300">Go / Wails / Svelte</strong>
          keeps the sensitive path on-device with a small, auditable surface.
        </p>
      </article>
    </div>
  </section>

  <section class="mb-12 border-t border-silo-border pt-8 font-sans" aria-labelledby="founder-heading">
    <h2
      id="founder-heading"
      class="mb-2 font-mono text-xs font-semibold uppercase tracking-widest text-silo-muted"
    >
      Founder
    </h2>
    <p class="text-sm leading-relaxed text-zinc-300">
      SiloRedact is developed by cybersecurity professional
      <strong class="text-white">Alexander Mckain</strong>. The tool grew from a personal need for
      <strong class="text-zinc-100">secure Local-First Architecture</strong>
      —all redaction and processing on-device—before sharing logs and tickets without trusting third-party
      clouds with raw sensitive data.
    </p>
    <div class="mt-4 flex flex-wrap items-center gap-2 text-xs">
      {#if appVersion}
        <span class="rounded bg-zinc-800 px-2 py-0.5 font-mono text-zinc-400">v{appVersion}</span>
      {/if}
      <span
        class="rounded border border-silo-accent/50 bg-silo-accent/10 px-2 py-0.5 font-medium text-silo-accent"
      >
        Open Source
      </span>
      <a
        href={githubUrl}
        target="_blank"
        rel="noopener noreferrer"
        class="text-silo-accent underline decoration-silo-accent/40 underline-offset-2 hover:text-cyan-300"
      >
        Source on GitHub
      </a>
    </div>
    <p class="mt-3 text-xs italic text-silo-muted">The Clipboard Firewall for the AI Era.</p>
    <div class="mt-4 flex flex-wrap gap-x-4 gap-y-2 text-sm">
      <a href={githubUrl} target="_blank" rel="noopener noreferrer" class="text-silo-accent hover:underline">
        GitHub
      </a>
      <a href={'mailto:' + contactEmail} class="text-silo-accent hover:underline">Contact</a>
    </div>
    <p class="mt-1 font-mono text-xs text-zinc-500">
      <a href={'mailto:' + contactEmail} class="text-zinc-400 hover:text-zinc-300">{contactEmail}</a>
    </p>
  </section>

  <!-- Legal archive -->
  <div class="border-t border-silo-border pt-10">
    <h2 class="mb-6 font-mono text-sm font-semibold uppercase tracking-widest text-zinc-500">
      Legal archive
    </h2>

    <section id="eula" class="mb-14 scroll-mt-24 rounded-lg border border-silo-border bg-zinc-900/30 p-5 sm:p-8">
      <header class="mb-6 border-b border-silo-border pb-4">
        <p class="font-mono text-xs uppercase tracking-wider text-silo-accent">Document</p>
        <h2 class="font-sans text-xl font-semibold text-white">End User License Agreement (EULA)</h2>
        <p class="mt-1 font-mono text-xs text-zinc-500">Effective: January 1, 2026 · SiloRedact software</p>
      </header>
      <div class={legalProse}>
        <p>
          This End User License Agreement (“<strong>Agreement</strong>”) is a legal agreement between you
          (“<strong>Licensee</strong>,” “<strong>you</strong>”) and Alexander Mckain (“<strong>Licensor</strong>”)
          for the SiloRedact software and related documentation (“<strong>Software</strong>”). By installing or
          using the Software, you agree to these terms. If you do not agree, do not use the Software.
        </p>
        <h3>1. License grant</h3>
        <p>
          Subject to this Agreement, Licensor grants you a limited, non-exclusive, non-transferable license to install
          and run the Software on devices you control, solely for your internal business or personal use in DevOps,
          cybersecurity, and IT workflows. SiloRedact is distributed as open-source software; see the project repository
          for the applicable open-source license terms.
        </p>
        <h3>2. DevOps &amp; security context</h3>
        <p>
          The Software is intended to assist with <strong>pre-paste redaction</strong> of logs, tickets, and
          similar material. You remain solely responsible for your security posture, compliance obligations
          (including SOC 2, HIPAA, PCI, GDPR, or sector-specific rules where applicable), and for any data you
          choose to transmit after redaction.
        </p>
        <h3>3. Verify your results</h3>
        <p>
          Redaction is pattern- and rule-driven and may not detect every sensitive token, secret, or identifier.
          <strong
            >You must independently verify that output is fit for your purpose before sharing, pasting, or
            storing it.</strong
          >
          Licensor does not warrant that redaction is complete, error-free, or suitable for any particular
          compliance regime. False negatives and false positives may occur.
        </p>
        <p>
          You acknowledge that SiloRedact is a tool to assist in data privacy and that the final responsibility for
          verifying the success of any redaction lies solely with you. We do not warrant that the Software will be
          error-free or that all sensitive data will be identified.
        </p>
        <h3>4. Local-First Architecture (core engine)</h3>
        <p>
          The <strong>core redaction engine</strong> implements <strong>Local-First Architecture</strong>:
          <strong
            >all file redaction and processing occur 100% on your local machine. Your data never leaves your
            device</strong
          >
          for that workload, and no redaction processing is performed in Licensor’s cloud solely by virtue of using those
          core features. Platform updates (including Microsoft Store distribution) may involve separate system-level
          network activity disclosed by your operating system. You acknowledge that your operating system, other
          applications, and destinations you paste into may still transmit data over networks outside Licensor’s control.
        </p>
        <h3>5. Restrictions</h3>
        <p>You shall not: (a) reverse engineer, decompile, or disassemble the Software except where mandatory law
          permits; (b) rent, lease, sell, sublicense, or timeshare the Software; (c) remove proprietary notices; (d)
          use the Software to violate law or third-party rights; (e) use the Software to build a competing
          redaction service that copies non-public aspects of the Software; or (f) interfere with license or
          integrity controls.</p>
        <h3>6. Ownership</h3>
        <p>
          Licensor retains all rights not expressly granted. The Software is licensed, not sold.
        </p>
        <h3>7. Updates</h3>
        <p>
          Licensor may provide updates or patches. Updates may be subject to additional terms presented at
          installation. If you obtain the Software from the Microsoft Store, Windows may also deliver updates as
          described in Microsoft’s terms.
        </p>
        <h3>8. Termination</h3>
        <p>
          Licensor may terminate this Agreement if you materially breach it. Upon termination, you must cease use
          and destroy copies in your possession.
        </p>
        <h3>9. Disclaimer of warranties</h3>
        <p>
          THE SOFTWARE IS PROVIDED “AS IS” AND “AS AVAILABLE.” TO THE MAXIMUM EXTENT PERMITTED BY LAW, LICENSOR
          DISCLAIMS ALL WARRANTIES, WHETHER EXPRESS, IMPLIED, OR STATUTORY, INCLUDING MERCHANTABILITY, FITNESS
          FOR A PARTICULAR PURPOSE, AND NON-INFRINGEMENT.
        </p>
        <h3>10. Limitation of liability</h3>
        <p>
          TO THE MAXIMUM EXTENT PERMITTED BY LAW, LICENSOR’S TOTAL LIABILITY ARISING OUT OF OR RELATED TO THIS
          AGREEMENT SHALL NOT EXCEED THE GREATER OF (A) THE FEES YOU PAID TO LICENSOR FOR THE SOFTWARE IN THE
          TWELVE (12) MONTHS BEFORE THE CLAIM, OR (B) FIFTY U.S. DOLLARS (USD $50). LICENSOR SHALL NOT BE LIABLE
          FOR ANY INDIRECT, INCIDENTAL, SPECIAL, CONSEQUENTIAL, OR PUNITIVE DAMAGES, OR FOR LOSS OF DATA, PROFITS,
          GOODWILL, OR SECURITY INCIDENTS ATTRIBUTABLE TO YOUR PASTE DESTINATIONS OR MISCONFIGURATION.
        </p>
        <h3>11. Export &amp; sanctions</h3>
        <p>
          You represent that you are not prohibited from receiving the Software under applicable export control
          or sanctions laws.
        </p>
        <h3>12. Governing law &amp; venue</h3>
        <p>
          This Agreement is governed by the laws of the <strong>State of California</strong>, without regard to
          conflict-of-law principles. Exclusive venue for disputes shall be the state and federal courts located in
          California, and you consent to personal jurisdiction there.
        </p>
        <h3>13. Entire agreement</h3>
        <p>
          This Agreement constitutes the entire agreement regarding the Software and supersedes prior
          understandings on the subject matter.
        </p>
      </div>
    </section>

    <section
      id="privacy"
      class="mb-14 scroll-mt-24 rounded-lg border border-silo-border bg-zinc-900/30 p-5 sm:p-8"
    >
      <header class="mb-6 border-b border-silo-border pb-4">
        <p class="font-mono text-xs uppercase tracking-wider text-silo-accent">Document</p>
        <h2 class="font-sans text-xl font-semibold text-white">Privacy Policy — Zero-Data posture</h2>
        <p class="mt-1 font-mono text-xs text-zinc-500">Effective: January 1, 2026 · SiloRedact</p>
      </header>
      <div class={legalProse}>
        <p>
          SiloRedact is built for operators who treat clipboard and log data as toxic. This policy describes how
          Licensor approaches information in connection with the <strong>desktop Software</strong>. It is not a
          substitute for your own enterprise privacy program.
        </p>
        <h3>1. No collection of your redacted content</h3>
        <p>
          The Software’s <strong>core redaction workflow does not upload the text you paste into the Software to
          Licensor for processing.</strong> Licensor does not operate a “cloud redaction” pipeline that ingests
          your clipboard or workspace content as part of normal local use.
        </p>
        <h3>2. No telemetry on your text</h3>
        <p>
          Licensor does not embed analytics that exfiltrate the substance of your logs, tickets, secrets, or
          redacted output. If optional diagnostics or crash reporting is ever offered, it will be
          <strong>clearly disclosed and opt-in</strong>, and will not be enabled here by default in a way that
          sends your redaction payload to Licensor without your knowledge.
        </p>
        <h3>3. No cloud-sync of your workspace</h3>
        <p>
          The Software does not sync your redaction workspace, rules, or buffers to Licensor-controlled cloud
          storage as a default “always-on” feature. Any backup or sync you configure (e.g., OS-level cloud
          drives) is outside this Software’s design perimeter.
        </p>
        <h3>4. Local-First processing</h3>
        <p>
          SiloRedact’s <strong>Local-First Architecture</strong> means core redaction and file processing stay on your
          device; your redaction buffer and file contents are not uploaded for cloud scrubbing as part of normal local
          use. The open-source desktop app does not require activation or payment flows.
        </p>
        <p>
          <strong>Microsoft Store:</strong> If you install via the Microsoft Store, Windows may occasionally connect to
          Microsoft servers for app updates. This is a system-level process and does not involve your processed data or
          redaction buffers.
        </p>
        <h3>5. What may be processed elsewhere</h3>
        <p>
          When you voluntarily use other network features (e.g., opening outbound links, CDNs, or update servers beyond
          the transactional flows above), providers may process connection or delivery metadata as described by their
          policies—not your redaction buffer or workspace content.
        </p>
        <h3>6. Communications</h3>
        <p>
          If you email <strong>support@siloredact.com</strong>, Licensor receives the content you send in that
          message.
        </p>
        <h3>7. Security</h3>
        <p>
          No method of transmission or storage is 100% secure. You are responsible for device hardening, disk
          encryption, and safe handling after redaction.
        </p>
        <h3>8. Children</h3>
        <p>The Software is not directed to children under 13.</p>
        <h3>9. Changes</h3>
        <p>
          Licensor may update this policy. Material changes will be indicated by revising the effective date at the
          top of this document.
        </p>
        <h3>10. Contact</h3>
        <p>
          Questions: <a href="mailto:support@siloredact.com">support@siloredact.com</a>
        </p>
      </div>
    </section>

    <section id="terms" class="mb-8 scroll-mt-24 rounded-lg border border-silo-border bg-zinc-900/30 p-5 sm:p-8">
      <header class="mb-6 border-b border-silo-border pb-4">
        <p class="font-mono text-xs uppercase tracking-wider text-silo-accent">Document</p>
        <h2 class="font-sans text-xl font-semibold text-white">Terms of Service</h2>
        <p class="mt-1 font-mono text-xs text-zinc-500">Effective: January 1, 2026 · SiloRedact</p>
      </header>
      <div class={legalProse}>
        <p>
          These Terms of Service (“<strong>Terms</strong>”) govern access to and use of SiloRedact websites,
          downloads, and related services (“<strong>Services</strong>”) offered by Alexander
          Mckain (“<strong>we</strong>,” “<strong>us</strong>”). The Software itself is also subject to the EULA
          above.
        </p>
        <h3>1. Eligibility</h3>
        <p>
          You must have authority to agree to these Terms. If you use the Services on behalf of an organization,
          you represent that you have authority to bind that organization.
        </p>
        <h3>2. License &amp; acceptable use</h3>
        <p>
          You agree to use the Services only for lawful purposes. You shall not: (a) probe, scan, or test the
          vulnerability of our systems without authorization; (b) overload or disrupt the Services; (c) attempt
          unauthorized access; (d) use the Services to distribute malware; (e) misrepresent affiliation with us; or
          (f) use the Services to facilitate unauthorized access to third-party systems.
        </p>
        <h3>3. Open source</h3>
        <p>
          SiloRedact is distributed as open-source software. Source code, contribution guidelines, and license terms
          are published in the project repository linked from the app and website.
        </p>
        <h3>4. Prohibited high-risk uses</h3>
        <p>
          You may not rely on the Services as the sole safeguard for life-critical infrastructure, weapons
          systems, or environments where failure could cause death or serious bodily injury, except where
          expressly agreed in a separate written contract.
        </p>
        <h3>5. Intellectual property</h3>
        <p>
          Content, branding, and documentation are owned by Licensor or licensors. No rights are granted except
          as expressly stated.
        </p>
        <h3>6. Third-party services</h3>
        <p>
          Payments, hosting, or email delivery may use third parties. Their terms apply to the extent you interact
          with them.
        </p>
        <h3>7. Disclaimers</h3>
        <p>
          SERVICES ARE PROVIDED “AS IS.” TO THE MAXIMUM EXTENT PERMITTED BY LAW, WE DISCLAIM WARRANTIES OF
          MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, AND NON-INFRINGEMENT.
        </p>
        <h3>8. Redaction verification</h3>
        <p>
          You acknowledge that SiloRedact is a tool to assist in data privacy and that the final responsibility for
          verifying the success of any redaction lies solely with the user. We do not warrant that the services will be
          error-free or that all sensitive data will be identified.
        </p>
        <h3>9. Limitation of liability</h3>
        <p>
          TO THE MAXIMUM EXTENT PERMITTED BY LAW, OUR TOTAL LIABILITY FOR ANY CLAIM RELATING TO THE SERVICES SHALL
          NOT EXCEED THE GREATER OF (A) AMOUNTS YOU PAID US IN THE TWELVE (12) MONTHS BEFORE THE CLAIM, OR (B) USD
          $50. WE ARE NOT LIABLE FOR INDIRECT, INCIDENTAL, SPECIAL, CONSEQUENTIAL, OR PUNITIVE DAMAGES, OR LOSS OF
          DATA OR PROFITS.
        </p>
        <h3>10. Indemnity</h3>
        <p>
          You will defend and indemnify us against claims arising from your misuse of the Services, your violation
          of these Terms, or your violation of third-party rights.
        </p>
        <h3>11. Governing law &amp; venue</h3>
        <p>
          These Terms are governed by the laws of the <strong>State of California</strong>, without regard to
          conflict-of-law principles. Exclusive venue for disputes shall be the state and federal courts located in
          California, and you consent to personal jurisdiction there.
        </p>
        <h3>12. Changes</h3>
        <p>
          We may modify these Terms. Continued use after changes become effective constitutes acceptance of the
          revised Terms.
        </p>
        <h3>13. Contact</h3>
        <p>
          <a href="mailto:support@siloredact.com">support@siloredact.com</a>
        </p>
      </div>
    </section>
  </div>
</div>
